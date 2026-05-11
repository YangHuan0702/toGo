package repository

import (
	"errors"
	"strings"
	"sync"
	"toGO/common"
	"toGO/contr/req"
	"toGO/domain"
	"toGO/domain/bo"
)

type PlanTodoRepositoryInterface interface {
	CreatePlan(domain.Plan) (int64, common.ApiException)
	CreateTodo(todo []domain.Todo) common.ApiException
	DeletePlan(id int64) common.ApiException
	DeleteTodo(id int64) common.ApiException
	GetPlanForId(id int64) *domain.Plan
	GetTodoForPlanId(id int64) []domain.Todo
	PageList(request req.TodoPageRequest) common.PageResp
}

var planStore = struct {
	sync.Mutex
	nextPlanID int64
	nextTodoID int64
	plans      []domain.Plan
	todos      []domain.Todo
}{
	nextPlanID: 4,
	nextTodoID: 6,
	plans: []domain.Plan{
		{Id: 1, UserName: "前端", Title: "完善创建计划弹窗", Content: "补齐校验、待办列表和提交反馈。", PlanStartDate: "2026-05-08", PlanFinishDate: "2026-05-15", PlanTag: "前端,联调"},
		{Id: 2, UserName: "后端", Title: "联调分页查询接口", Content: "统一分页字段和错误提示。", PlanStartDate: "2026-05-09", PlanFinishDate: "2026-05-12", PlanTag: "后端,P1"},
		{Id: 3, UserName: "测试", Title: "移动端布局检查", Content: "重点检查菜单、表格和弹窗宽度。", PlanStartDate: "2026-05-01", PlanFinishDate: "2026-05-07", PlanTag: "测试,复盘"},
	},
	todos: []domain.Todo{
		{Id: 1, PlanId: 1, Title: "基础表单校验", Remark: "必填和日期格式", Content: "为计划创建表单补充完整校验。", StartDate: "2026-05-08", EndDate: "2026-05-10"},
		{Id: 2, PlanId: 1, Title: "待办明细录入", Remark: "支持多条", Content: "提交时携带 items 数组。", StartDate: "2026-05-10", EndDate: "2026-05-15"},
		{Id: 3, PlanId: 2, Title: "分页假数据", Remark: "不连接数据库", Content: "根据前端字段返回 planTitle 和 todoContent。", StartDate: "2026-05-09", EndDate: "2026-05-11"},
		{Id: 4, PlanId: 2, Title: "创建计划接口", Remark: "内存追加", Content: "支持 /plan/create 保存计划和待办。", StartDate: "2026-05-11", EndDate: "2026-05-12"},
		{Id: 5, PlanId: 3, Title: "断点检查", Remark: "手机和桌面", Content: "检查列表页、菜单和弹窗布局。", StartDate: "2026-05-01", EndDate: "2026-05-07"},
	},
}

func GetPlanRepository() *PlanTodoRepository {
	return &PlanTodoRepository{}
}

type PlanTodoRepository struct {
	PlanTodoRepositoryInterface
}

func (p *PlanTodoRepository) CreatePlan(plan domain.Plan) (int64, common.ApiException) {
	planStore.Lock()
	defer planStore.Unlock()

	plan.Id = planStore.nextPlanID
	planStore.nextPlanID++
	planStore.plans = append(planStore.plans, plan)
	return plan.Id, nil
}

func (p *PlanTodoRepository) CreatePlanWithTodos(request req.PlanCreateRequest) (domain.Plan, error) {
	if strings.TrimSpace(request.Title) == "" {
		return domain.Plan{}, errors.New("计划标题不能为空")
	}
	if strings.TrimSpace(request.UserName) == "" {
		return domain.Plan{}, errors.New("负责人不能为空")
	}
	if len(request.Items) == 0 {
		return domain.Plan{}, errors.New("请至少添加一个待办")
	}

	planStore.Lock()
	defer planStore.Unlock()

	plan := domain.Plan{
		Id:             planStore.nextPlanID,
		UserName:       request.UserName,
		Title:          request.Title,
		Content:        request.Content,
		PlanFinishDate: request.PlanFinishDate,
		PlanStartDate:  request.PlanStartDate,
		PlanTag:        request.PlanTag,
		Items:          make([]domain.Todo, 0, len(request.Items)),
	}
	planStore.nextPlanID++
	planStore.plans = append(planStore.plans, plan)

	for _, item := range request.Items {
		todo := domain.Todo{
			Id:        planStore.nextTodoID,
			PlanId:    plan.Id,
			Title:     item.Title,
			Remark:    item.Remark,
			Content:   item.Content,
			StartDate: item.StartDate,
			EndDate:   item.EndDate,
		}
		planStore.nextTodoID++
		planStore.todos = append(planStore.todos, todo)
		plan.Items = append(plan.Items, todo)
	}

	return plan, nil
}

func (p *PlanTodoRepository) CreateTodo(todo []domain.Todo) common.ApiException {
	planStore.Lock()
	defer planStore.Unlock()

	for idx := range todo {
		todo[idx].Id = planStore.nextTodoID
		planStore.nextTodoID++
		planStore.todos = append(planStore.todos, todo[idx])
	}
	return nil
}

func (p *PlanTodoRepository) DeletePlan(id int64) common.ApiException {
	planStore.Lock()
	defer planStore.Unlock()

	for idx := range planStore.plans {
		if planStore.plans[idx].Id == id {
			planStore.plans = append(planStore.plans[:idx], planStore.plans[idx+1:]...)
			return nil
		}
	}
	return common.ExceptionRespMap[common.NotFindPlan]
}

func (p *PlanTodoRepository) DeleteTodo(id int64) common.ApiException {
	planStore.Lock()
	defer planStore.Unlock()

	for idx := range planStore.todos {
		if planStore.todos[idx].Id == id {
			planStore.todos = append(planStore.todos[:idx], planStore.todos[idx+1:]...)
			return nil
		}
	}
	return nil
}

func (p *PlanTodoRepository) GetPlanForId(id int64) *domain.Plan {
	planStore.Lock()
	defer planStore.Unlock()

	for idx := range planStore.plans {
		if planStore.plans[idx].Id == id {
			plan := planStore.plans[idx]
			plan.Items = p.todosForPlanLocked(id)
			return &plan
		}
	}
	return nil
}

func (p *PlanTodoRepository) GetTodoForPlanId(id int64) []domain.Todo {
	planStore.Lock()
	defer planStore.Unlock()

	return p.todosForPlanLocked(id)
}

func (p *PlanTodoRepository) PageList(request req.TodoPageRequest) common.PageResp {
	planStore.Lock()
	defer planStore.Unlock()

	if request.CurPage <= 0 {
		request.CurPage = 1
	}
	if request.PageSize <= 0 {
		request.PageSize = 20
	}

	rows := make([]bo.PlanPageBo, 0)
	for _, plan := range planStore.plans {
		if !contains(plan.Title, request.PlanTitle) || !contains(plan.UserName, request.UserName) {
			continue
		}

		hasTodo := false
		for _, todo := range planStore.todos {
			if todo.PlanId != plan.Id || !contains(todo.Title, request.TodoTitle) {
				continue
			}
			hasTodo = true
			rows = append(rows, planPageRow(plan, todo))
		}

		if !hasTodo && request.TodoTitle == "" {
			rows = append(rows, planPageRow(plan, domain.Todo{}))
		}
	}

	count := int64(len(rows))
	start := (request.CurPage - 1) * request.PageSize
	if start > len(rows) {
		start = len(rows)
	}
	end := start + request.PageSize
	if end > len(rows) {
		end = len(rows)
	}

	return common.PageResp{
		Data:      rows[start:end],
		Count:     count,
		CurPage:   request.CurPage,
		PageSize:  request.PageSize,
		PageCount: count,
	}
}

func (p *PlanTodoRepository) todosForPlanLocked(planID int64) []domain.Todo {
	todos := make([]domain.Todo, 0)
	for _, todo := range planStore.todos {
		if todo.PlanId == planID {
			todos = append(todos, todo)
		}
	}
	return todos
}

func planPageRow(plan domain.Plan, todo domain.Todo) bo.PlanPageBo {
	return bo.PlanPageBo{
		PlanId:         plan.Id,
		PlanTitle:      plan.Title,
		UserName:       plan.UserName,
		PlanFinishDate: plan.PlanFinishDate,
		PlanStartDate:  plan.PlanStartDate,
		PlanTag:        plan.PlanTag,
		TodoTitle:      todo.Title,
		TodoRemark:     todo.Remark,
		TodoContent:    todo.Content,
		TodoStartDate:  todo.StartDate,
		TodoEndDate:    todo.EndDate,
	}
}

func contains(value string, keyword string) bool {
	return keyword == "" || strings.Contains(strings.ToLower(value), strings.ToLower(keyword))
}
