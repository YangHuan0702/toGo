package contr

import (
	"toGO/common"
	"toGO/contr/req"
	"toGO/repository"
)

func GetPlanTodoController() *PlanTodoController {
	return &PlanTodoController{service: repository.GetPlanRepository()}
}

type PlanTodoController struct {
	service *repository.PlanTodoRepository
}

func (contr *PlanTodoController) PageList(params req.TodoPageRequest) common.ToGoResponse {
	return common.Success(contr.service.PageList(params))
}

func (contr *PlanTodoController) PlanInfo(id int64) common.ToGoResponse {
	plan := contr.service.GetPlanForId(id)
	return common.Success(plan)
}

func (contr *PlanTodoController) PlanTodos(planId int64) common.ToGoResponse {
	return common.Success(contr.service.GetTodoForPlanId(planId))
}
