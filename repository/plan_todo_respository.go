package repository

import (
	"fmt"
	"toGO/common"
	"toGO/contr/req"
	"toGO/domain"
	"toGO/domain/bo"

	"gorm.io/gorm"
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

func GetPlanRepository() *PlanTodoRepository {
	return &PlanTodoRepository{db: GetDBConn()}
}

type PlanTodoRepository struct {
	db *gorm.DB
	PlanTodoRepositoryInterface
}

func (p *PlanTodoRepository) CreatePlan(plan domain.Plan) (int64, common.ApiException) {
	p.db.Create(&plan)
	return plan.Id, nil
}

func (p *PlanTodoRepository) CreateTodo(todo []domain.Todo) common.ApiException {
	p.db.CreateInBatches(todo, len(todo))
	return nil
}

func (p *PlanTodoRepository) DeletePlan(id int64) common.ApiException {
	plan := p.GetPlanForId(id)
	if plan == nil {
		return common.ExceptionRespMap[common.NotFindPlan]
	}

	p.db.Delete(&domain.Plan{}, id)
	return nil
}

func (p *PlanTodoRepository) DeleteTodo(id int64) common.ApiException {
	p.db.Delete(&domain.Todo{}, id)
	return nil
}

func (p *PlanTodoRepository) GetPlanForId(id int64) *domain.Plan {
	plan := &domain.Plan{}
	p.db.Where("id = ?", id).First(plan)
	return plan
}

func (p *PlanTodoRepository) GetTodoForPlanId(id int64) []domain.Todo {
	todos := make([]domain.Todo, 0)
	p.db.Where("plan_id = ?", id).Find(&todos)
	return todos
}

func (p *PlanTodoRepository) PageList(request req.TodoPageRequest) common.PageResp {

	sql :=
		`select p.id as plan_id,p.title as plan_title, p.user_name,
       p.plan_finish_date, p.plan_start_date,p.plan_tag,t.title as todo_title, t.remark as todo_remark, t.content as todo_content,t.start_date as todo_start_date, t.end_date as todo_end_date
from t_plan p left join t_todo t on p.id = t.plan_id where 1 = 1`

	if len(request.PlanTitle) > 0 {
		sql += fmt.Sprintf(" and p.plan_title = '%s'", request.PlanTitle)
	}
	if len(request.UserName) > 0 {
		sql += fmt.Sprintf(" and p.user_name = '%s'", request.UserName)
	}
	if len(request.TodoTitle) > 0 {
		sql += fmt.Sprintf(" and t.title = '%s'", request.TodoTitle)
	}

	rows := make([]bo.PlanPageBo, 0)
	p.db.Raw(sql).Limit(request.PageSize).Offset((request.CurPage - 1) * request.PageSize).Scan(&rows)
	var count int64 = 0
	p.db.Raw(sql).Count(&count)

	return common.PageResp{
		Data:      rows,
		Count:     count,
		CurPage:   request.CurPage,
		PageSize:  request.PageSize,
		PageCount: (count + 1) / int64(request.PageSize),
	}
}
