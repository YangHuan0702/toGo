package req

type MenuPageRequest struct {
	Name     string `json:"name"`
	PageNo   int    `json:"pageNo"`
	PageSize int    `json:"pageSize"`
}

type MenuSaveRequest struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Component string `json:"component"`
	ParentId  int64  `json:"parentId"`
}
