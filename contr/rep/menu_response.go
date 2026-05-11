package rep

import "toGO/domain"

type MenuListResponse struct {
	domain.Menu

	// 子菜单
	Children []*MenuListResponse `json:"children"`
}

func ConversionMenuToMenuListResp(menu *domain.Menu) *MenuListResponse {
	return &MenuListResponse{
		Menu:     *menu,
		Children: make([]*MenuListResponse, 0),
	}
}
