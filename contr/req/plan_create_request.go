package req

type PlanCreateRequest struct {
	UserName       string                  `json:"userName"`
	Title          string                  `json:"title"`
	Content        string                  `json:"content"`
	PlanFinishDate string                  `json:"planFinishDate"`
	PlanStartDate  string                  `json:"planStartDate"`
	PlanTag        string                  `json:"planTag"`
	Items          []PlanItemCreateRequest `json:"items"`
}

type PlanItemCreateRequest struct {
	Title     string `json:"title"`
	Remark    string `json:"remark"`
	Content   string `json:"content"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}
