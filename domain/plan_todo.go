package domain

type Plan struct {
	Id             int64  `gorm:"primary_key;auto_increment" json:"id"`
	UserName       string `gorm:"user_name" json:"userName"`
	Title          string `gorm:"title" json:"title"`
	Content        string `gorm:"content" json:"content"`
	PlanFinishDate string `gorm:"plan_finish_date" json:"planFinishDate"`
	PlanStartDate  string `gorm:"plan_start_date" json:"planStartDate"`
	PlanTag        string `gorm:"plan_tag" json:"planTag"`
}

func (Plan) TableName() string {
	return "t_plan"
}

type Todo struct {
	Id        int64  `gorm:"primary_key;auto_increment" json:"id"`
	PlanId    int64  `gorm:"plan_id" json:"planId"`
	Title     string `gorm:"title" json:"title"`
	Remark    string `gorm:"remark" json:"remark"`
	Content   string `gorm:"content" json:"content"`
	StartDate string `gorm:"start_date" json:"startDate"`
	EndDate   string `gorm:"end_date" json:"endDate"`
}

func (Todo) TableName() string {
	return "t_todo"
}
