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
	nextPlanID: 1,
	nextTodoID: 1,
	plans:      []domain.Plan{},
	todos:      []domain.Todo{},
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
