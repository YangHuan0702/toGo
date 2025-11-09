package repository

import (
	"toGO/domain"

	"toGO/common"

	"gorm.io/gorm"
)

type MenuRepositoryInterface interface {

	// CreateMenu 创建菜单
	CreateMenu(domain.Menu) (int64, common.ApiException)

	// UpdateMenu 修改菜单
	UpdateMenu(domain.Menu) common.ApiException

	// DeleteMenu 删除菜单
	DeleteMenu(int64) (common.ApiException, *domain.Menu)

	// GetById 通过ID查询菜单
	GetById(int64) *domain.Menu

	// PageList 通过菜单名称进行分类查询
	PageList(string, int, int) common.PageResp
}

type MenuRepository struct {
	db *gorm.DB
	MenuRepositoryInterface
}

func (rep *MenuRepository) CreateMenu(menu *domain.Menu) (int64, common.ApiException) {
	if menu == nil {
		return -1, common.ExceptionRespMap[common.NilParams]
	}

	// 父级菜单是否存在
	if menu.ParentId != 0 {
		parent := rep.GetById(menu.ParentId)
		if parent == nil {
			return -1, common.ExceptionRespMap[common.NotFindParentMenu]
		}
	}

	rep.db.Create(menu)
	return menu.Id, nil
}

func (rep *MenuRepository) UpdateMenu(menu domain.Menu) common.ApiException {
	if menu.ParentId != 0 {
		parent := rep.GetById(menu.ParentId)
		if parent == nil {
			return common.ExceptionRespMap[common.NotFindParentMenu]
		}
	}

	self := rep.GetById(menu.Id)
	if self == nil {
		return common.ExceptionRespMap[common.NotFindMenu]
	}

	rep.db.Save(&menu)
	return nil
}

func (rep *MenuRepository) DeleteMenu(id int64) (common.ApiException, *domain.Menu) {
	target := rep.GetById(id)

	if target == nil {
		return common.ExceptionRespMap[common.NotFindMenu], nil
	}

	rep.db.Delete(target)
	return nil, target
}

func (rep *MenuRepository) GetById(id int64) *domain.Menu {
	menu := &domain.Menu{}
	rep.db.Where("id = ?", id).First(&menu)
	return menu
}

func (rep *MenuRepository) PageList(menuName string, pageNo int, pageSize int) common.PageResp {
	var count int64 = 0
	var users []domain.Menu
	if menuName != "" {
		rep.db.Model(&domain.Menu{}).Where("name = ?", menuName).Count(&count)
		rep.db.Limit(pageSize).Offset((pageNo-1)*pageSize).Where("name = ?", menuName).Find(&users)
	} else {
		rep.db.Model(&domain.Menu{}).Count(&count)
		rep.db.Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&users)
	}
	return common.PageResp{Data: users, Count: count}
}
