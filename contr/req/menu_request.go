package req

type MenuPageRequest struct {
	Name     string `json:"name"`
	PageNo   int    `json:"pageNo"`
	PageSize int    `json:"pageSize"`
}
