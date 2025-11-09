package contr

import (
	"toGO/common"
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
