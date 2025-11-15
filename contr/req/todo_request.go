package req

type TodoPageRequest struct {
	CurPage   int    `json:"curPage"`
	PageSize  int    `json:"pageSize"`
	PlanTitle string `json:"planTitle"`
	TodoTitle string `json:"todoTitle"`
	UserName  string `json:"userName"`
}
