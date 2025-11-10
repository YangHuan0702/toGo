package contr

import (
	"toGO/common"
	"toGO/contr/rep"
	"toGO/contr/req"
	"toGO/repository"
)

func GetMenuController() MenuController {
	menuController := MenuController{menuService: repository.GetMenuRepository()}
	return menuController
}

type MenuController struct {
	menuService repository.MenuRepository
}

func (contr *MenuController) PageList(req req.MenuPageRequest) common.ToGoResponse {
	return common.Success(contr.menuService.PageList(req.Name, req.PageNo, req.PageSize))
}

func (contr *MenuController) List(name string) common.ToGoResponse {
	menus := contr.menuService.ListMenuForName(name)

	returnOfParent := make([]*rep.MenuListResponse, 0)

	menuMap := make(map[int64]*rep.MenuListResponse)
	for idx := range menus {
		menu := &menus[idx]
		resp := rep.ConversionMenuToMenuListResp(menu)
		if menu.ParentId == 0 {
			returnOfParent = append(returnOfParent, resp)
		}
		menuMap[menu.Id] = resp
	}

	for idx := range menus {
		menu := &menus[idx]
		if menu.ParentId != 0 && menuMap[menu.ParentId] != nil {
			menuMap[menu.ParentId].Children = append(menuMap[menu.ParentId].Children, menuMap[menu.Id])
		}
	}
	return common.Success(returnOfParent)
}
