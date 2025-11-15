package bo

type PlanPageBo struct {
	PlanId         int64  `json:"planId" gorm:"plan_id"`
	PlanTitle      string `json:"planTitle" gorm:"plan_title"`
	UserName       string `json:"userName" gorm:"user_name"`
	PlanFinishDate string `json:"planFinishDate" gorm:"plan_finish_date"`
	PlanStartDate  string `json:"planStartDate" gorm:"plan_start_date"`
	PlanTag        string `json:"planTag" gorm:"plan_tag"`

	TodoTitle     string `json:"todoTitle" gorm:"todo_title"`
	TodoRemark    string `json:"todoRemark" gorm:"todo_remark"`
	TodoContent   string `json:"TodoContent" gorm:"todo_content"`
	TodoStartDate string `json:"todoStartDate" grom:"todo_start_date"`
	TodoEndDate   string `json:"todoEndDate" grom:"todo_end_date"`
}
