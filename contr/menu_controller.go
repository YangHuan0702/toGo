package contr

import (
	"toGO/common"
	"toGO/contr/req"
	"toGO/repository"
)

type MenuController struct {
	menuService repository.MenuRepository
}

func (contr *MenuController) PageList(req req.MenuPageRequest) common.ToGoResponse {
	return common.Success(contr.menuService.PageList(req.Name, req.PageNo, req.PageSize))

}
