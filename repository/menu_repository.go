package repository

import (
	"toGO/domain"

	"toGO/common"

	"gorm.io/gorm"
)

type MenuRepositoryInterface interface {

	// 创建菜单
	createMenu(domain.Menu) (int64, common.ApiException)

	// 修改菜单
	updateMenu(domain.Menu) domain.Menu

	// 删除菜单
	deleteMenu(int64) domain.Menu

	// 通过ID查询菜单
	getById(int64) *domain.Menu

	// 通过菜单名称进行分类查询
	pageList(string, int, int) []domain.Menu
}

type MenuRepository struct {
	db *gorm.DB
	MenuRepositoryInterface
}

func (rep *MenuRepository) createMenu(menu *domain.Menu) (int64, common.ApiException) {
	if menu == nil {
		return -1, common.ExceptionRespMap[common.NilParams]
	}

	// 父级菜单是否存在
	if menu.ParentId != 0 {
		parent := rep.getById(menu.ParentId)
		if parent == nil {
			return -1, common.ExceptionRespMap[common.NotFindParentMenu]
		}
	}

	rep.db.Create(menu)
	return menu.Id, nil
}

func (*MenuRepository) updateMenu(domain.Menu) domain.Menu {

}

func (*MenuRepository) deleteMenu(int64) domain.Menu {

}

func (*MenuRepository) getById(id int64) *domain.Menu {

}

func (*MenuRepository) pageList(string, int, int) []domain.Menu {

}
