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
	updateMenu(domain.Menu) common.ApiException

	// 删除菜单
	deleteMenu(int64) (common.ApiException, *domain.Menu)

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

func (rep *MenuRepository) updateMenu(menu domain.Menu) common.ApiException {
	if menu.ParentId != 0 {
		parent := rep.getById(menu.ParentId)
		if parent == nil {
			return common.ExceptionRespMap[common.NotFindParentMenu]
		}
	}

	self := rep.getById(menu.Id)
	if self == nil {
		return common.ExceptionRespMap[common.NotFindMenu]
	}

	rep.db.Save(&menu)
	return nil
}

func (rep *MenuRepository) deleteMenu(id int64) (common.ApiException, *domain.Menu) {
	target := rep.getById(id)

	if target == nil {
		return common.ExceptionRespMap[common.NotFindMenu], nil
	}

	rep.db.Delete(target)
	return nil, target
}

func (rep *MenuRepository) getById(id int64) *domain.Menu {
	menu := &domain.Menu{}
	rep.db.Where("id = ?", id).First(&menu)
	return menu
}

func (rep *MenuRepository) pageList(menuName string, pageNo int, pageSize int) []domain.Menu {
	var count int64 = 0
	var users []domain.Menu
	if menuName != "" {
		rep.db.Model(&domain.Menu{}).Where("name = ?", menuName).Count(&count)
		rep.db.Limit(pageSize).Offset((pageNo-1)*pageSize).Where("name = ?", menuName).Find(&users)
	} else {
		rep.db.Model(&domain.Menu{}).Count(&count)
		rep.db.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&users)
	}
	return users
}
