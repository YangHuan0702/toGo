package repository

import (
	"strings"
	"toGO/common"
	"toGO/domain"
)

type MenuRepositoryInterface interface {
	CreateMenu(*domain.Menu) (int64, common.ApiException)
	UpdateMenu(domain.Menu) common.ApiException
	DeleteMenu(int64) (common.ApiException, *domain.Menu)
	GetById(int64) *domain.Menu
	PageList(string, int, int) common.PageResp
	ListMenuForName(string) []domain.Menu
}

var menuRows = []domain.Menu{
	{Id: 1, Name: "待办计划管理", Path: "/todo", Component: "Layout", ParentId: 0},
	{Id: 2, Name: "计划概览", Path: "/todo/overview", Component: "todo/overview", ParentId: 1},
	{Id: 3, Name: "我的计划", Path: "/todo/list", Component: "todo/list", ParentId: 1},
	{Id: 4, Name: "任务看板", Path: "/todo/board", Component: "todo/board", ParentId: 1},
	{Id: 5, Name: "计划日历", Path: "/todo/calendar", Component: "todo/calendar", ParentId: 1},
	{Id: 6, Name: "复盘回顾", Path: "/todo/review", Component: "todo/review", ParentId: 1},
	{Id: 7, Name: "学习编排", Path: "/learning", Component: "Layout", ParentId: 0},
	{Id: 8, Name: "学习路径", Path: "/learning/path", Component: "learning/path", ParentId: 7},
	{Id: 9, Name: "学习计划", Path: "/learning/schedule", Component: "learning/schedule", ParentId: 7},
	{Id: 10, Name: "学习资料", Path: "/learning/materials", Component: "learning/materials", ParentId: 7},
	{Id: 11, Name: "复习记录", Path: "/learning/review", Component: "learning/review", ParentId: 7},
}

func GetMenuRepository() MenuRepository {
	return MenuRepository{}
}

type MenuRepository struct {
	MenuRepositoryInterface
}

func (rep *MenuRepository) CreateMenu(menu *domain.Menu) (int64, common.ApiException) {
	if menu == nil {
		return -1, common.ExceptionRespMap[common.NilParams]
	}

	var maxID int64
	for _, row := range menuRows {
		if row.Id > maxID {
			maxID = row.Id
		}
	}
	menu.Id = maxID + 1
	menuRows = append(menuRows, *menu)
	return menu.Id, nil
}

func (rep *MenuRepository) UpdateMenu(menu domain.Menu) common.ApiException {
	for idx := range menuRows {
		if menuRows[idx].Id == menu.Id {
			menuRows[idx] = menu
			return nil
		}
	}
	return common.ExceptionRespMap[common.NotFindMenu]
}

func (rep *MenuRepository) DeleteMenu(id int64) (common.ApiException, *domain.Menu) {
	for idx := range menuRows {
		if menuRows[idx].Id == id {
			menu := menuRows[idx]
			menuRows = append(menuRows[:idx], menuRows[idx+1:]...)
			return nil, &menu
		}
	}
	return common.ExceptionRespMap[common.NotFindMenu], nil
}

func (rep *MenuRepository) GetById(id int64) *domain.Menu {
	for idx := range menuRows {
		if menuRows[idx].Id == id {
			menu := menuRows[idx]
			return &menu
		}
	}
	return nil
}

func (rep *MenuRepository) PageList(menuName string, pageNo int, pageSize int) common.PageResp {
	if pageNo <= 0 {
		pageNo = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	rows := rep.ListMenuForName(menuName)
	count := int64(len(rows))
	start := (pageNo - 1) * pageSize
	if start > len(rows) {
		start = len(rows)
	}
	end := start + pageSize
	if end > len(rows) {
		end = len(rows)
	}

	return common.PageResp{
		Data:      rows[start:end],
		Count:     count,
		CurPage:   pageNo,
		PageSize:  pageSize,
		PageCount: count,
	}
}

func (rep *MenuRepository) ListMenuForName(name string) []domain.Menu {
	menus := make([]domain.Menu, 0)
	for _, menu := range menuRows {
		if name == "" || strings.Contains(strings.ToLower(menu.Name), strings.ToLower(name)) {
			menus = append(menus, menu)
		}
	}
	return menus
}
