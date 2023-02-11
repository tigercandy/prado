package system

import "github.com/tigercandy/prado/internal/models/system"

type sysMenuService struct {
}

var (
	SysMenuService = new(sysMenuService)
	menuModel      = new(system.MenuModel)
)

func (sms *sysMenuService) SetMenuLabel() (m []system.MenuLabel, err error) {
	menuList, err := menuModel.Get()
	m = make([]system.MenuLabel, 0)
	for i := 0; i < len(menuList); i++ {
		if menuList[i].ParentId != 0 {
			continue
		}
		e := system.MenuLabel{}
		e.Id = menuList[i].MenuId
		e.Label = menuList[i].Title
		menusInfo := sms.DiguiMenuLabel(&menuList, e)
		m = append(m, menusInfo)
	}
	return
}

func (sms *sysMenuService) DiguiMenuLabel(menuList *[]system.MenuModel, menu system.MenuLabel) system.MenuLabel {
	list := *menuList

	min := make([]system.MenuLabel, 0)
	for j := 0; j < len(list); j++ {
		if menu.Id != list[j].ParentId {
			continue
		}
		m := system.MenuLabel{}
		m.Id = list[j].MenuId
		m.Label = list[j].Title
		m.Children = []system.MenuLabel{}
		if list[j].MenuType != "F" {
			ms := sms.DiguiMenuLabel(menuList, m)
			min = append(min, ms)
		} else {
			min = append(min, m)
		}
	}
	menu.Children = min
	return menu
}

func (sms *sysMenuService) SetMenuRole(roleName string) (m []system.MenuModel, err error) {
	menuList, err := menuModel.GetByRoleName(roleName)

	m = make([]system.MenuModel, 0)
	for i := 0; i < len(menuList); i++ {
		if menuList[i].ParentId != 0 {
			continue
		}
		menusInfo := sms.DiguiMenu(&menuList, menuList[i])
		m = append(m, menusInfo)
	}
	return
}

func (sms *sysMenuService) DiguiMenu(menuList *[]system.MenuModel, menu system.MenuModel) system.MenuModel {
	list := *menuList
	min := make([]system.MenuModel, 0)
	for j := 0; j < len(list); j++ {
		if menu.MenuId != list[j].ParentId {
			continue
		}
		m := system.MenuModel{}
		m.MenuId = list[j].MenuId
		m.MenuName = list[j].MenuName
		m.Title = list[j].Title
		m.Icon = list[j].Icon
		m.Path = list[j].Path
		m.MenuType = list[j].MenuType
		m.Action = list[j].Action
		m.Permission = list[j].Permission
		m.ParentId = list[j].ParentId
		m.NoCache = list[j].NoCache
		m.BreadCrumb = list[j].BreadCrumb
		m.Component = list[j].Component
		m.Sort = list[j].Sort
		m.Visible = list[j].Visible
		m.CreatedAt = list[j].CreatedAt
		m.UpdatedAt = list[j].UpdatedAt
		m.Children = []system.MenuModel{}

		if m.MenuType != "F" {
			ms := sms.DiguiMenu(menuList, m)
			min = append(min, ms)
		} else {
			min = append(min, m)
		}
	}

	menu.Children = min
	return menu
}

func (sms *sysMenuService) SetMenu() (m []system.MenuModel, err error) {
	var menuModel = new(system.MenuModel)
	menuList, err := menuModel.GetPage()

	m = make([]system.MenuModel, 0)
	for i := 0; i < len(menuList); i++ {
		if menuList[i].ParentId != 0 {
			continue
		}
		menuInfo := sms.DiguiMenu(&menuList, menuList[i])

		m = append(m, menuInfo)
	}
	return
}
