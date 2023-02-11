package system

import (
	"github.com/tigercandy/prado/global/orm"
	"github.com/tigercandy/prado/pkg/utils"
)

type MenuModel struct {
	MenuId      int         `gorm:"primary_key;AUTO_INCREMENT" json:"menuId"`
	MenuName    string      `gorm:"type:varchar(128)" json:"menuName"`
	Title       string      `gorm:"type:varchar(64)" json:"title"`
	Icon        string      `gorm:"type:varchar(128)" json:"icon"`
	Path        string      `gorm:"type:varchar(128)" json:"path"`
	Paths       string      `gorm:"type:varchar(128)" json:"paths"`
	MenuType    string      `gorm:"type:varchar(1)" json:"menuType"`
	Action      string      `gorm:"type:varchar(16)" json:"action"`
	Permission  string      `gorm:"type:varchar(32)" json:"permission"`
	ParentId    int         `gorm:"type:int(11)" json:"parentId"`
	NoCache     bool        `gorm:"type:char(1)" json:"noCache"`
	BreadCrumb  string      `gorm:"type:varchar(255)" json:"breadCrumb"`
	Component   string      `gorm:"type:varchar(255)" json:"component"`
	Sort        int         `gorm:"type:int(4)" json:"sort"`
	Visible     string      `gorm:"type:char(1)" json:"visible"`
	CreatedUser string      `gorm:"type:varchar(128)" json:"createdUser"`
	UpdatedUser string      `gorm:"type:varchar(128)" json:"updatedUser"`
	IsFrame     string      `gorm:"type:int(1)" json:"isFrame"`
	Params      string      `gorm:"-" json:"params"`
	RoleId      int         `gorm:"-"`
	Children    []MenuModel `gorm:"-" json:"children"`
	IsSelect    bool        `gorm:"-" json:"is_select"`
	BaseModel
}

type Menus struct {
	MenuId      int         `gorm:"column:menu_id;primary_key;" json:"menuId"`
	MenuName    string      `gorm:"type:varchar(128)" json:"menuName"`
	Title       string      `gorm:"type:varchar(64)" json:"title"`
	Icon        string      `gorm:"type:varchar(128)" json:"icon"`
	Path        string      `gorm:"type:varchar(128)" json:"path"`
	MenuType    string      `gorm:"type:varchar(1)" json:"menuType"`
	Action      string      `gorm:"type:varchar(16)" json:"action"`
	Permission  string      `gorm:"type:varchar(32)" json:"permission"`
	ParentId    int         `gorm:"type:int(11)" json:"parentId"`
	NoCache     bool        `gorm:"type:char(1)" json:"noCache"`
	BreadCrumb  string      `gorm:"type:varchar(255)" json:"breadCrumb"`
	Component   string      `gorm:"type:varchar(255)" json:"component"`
	Sort        int         `gorm:"type:int(4)" json:"sort"`
	Visible     string      `gorm:"type:char(1)" json:"visible"`
	Children    []MenuModel `gorm:"-" json:"children"`
	CreatedUser string      `gorm:"type:varchar(128)" json:"createdUser"`
	UpdatedUser string      `gorm:"type:varchar(128)" json:"updatedUser"`
	Params      string      `gorm:"-" json:"params"`
	BaseModel
}

type MenuLabel struct {
	Id       int         `gorm:"-" json:"id"`
	Label    string      `gorm:"-" json:"label"`
	Children []MenuLabel `gorm:"-" json:"children"`
}

type MenuRole struct {
	Menus
	IsSelect bool `gorm:"-" json:"is_select"`
}

func (MenuModel) TableName() string {
	return "sys_menu"
}

func (e *MenuModel) GetMenuById() (MenuModel MenuModel, err error) {
	table := orm.Eloquent.Table(e.TableName())
	table = table.Where("menu_id = ? ", e.MenuId)
	if err = table.Find(&MenuModel).Error; err != nil {
		return
	}

	return
}

func (e *MenuModel) Get() (Menus []MenuModel, err error) {
	table := orm.Eloquent.Table(e.TableName())
	if e.MenuName != "" {
		table = table.Where("menu_name = ? ", e.MenuName)
	}
	if e.Path != "" {
		table = table.Where("path = ? ", e.Path)
	}
	if e.Action != "" {
		table = table.Where("action = ? ", e.Action)
	}
	if e.MenuType != "" {
		table = table.Where("menu_type = ? ", e.MenuType)
	}

	if err = table.Order("sort").Find(&Menus).Error; err != nil {
		return
	}

	return
}

func (e *MenuModel) Create() (id int, err error) {
	result := orm.Eloquent.Table(e.TableName()).Create(&e)
	if result.Error != nil {
		err = result.Error
		return
	}
	err = e.InitPaths(e)
	if err != nil {
		return
	}
	id = e.MenuId
	return
}

func (e *MenuModel) InitPaths(menu *MenuModel) (err error) {
	parentMenu := new(MenuModel)
	if int(menu.ParentId) != 0 {
		orm.Eloquent.Table(e.TableName()).Where("menu_id = ? ", menu.ParentId).First(parentMenu)
		if parentMenu.Paths == "" {
			return
		}
		menu.Paths = parentMenu.Paths + "/" + utils.Int2String(menu.MenuId)
	} else {
		menu.Paths = "/0/" + utils.Int2String(menu.MenuId)
	}
	orm.Eloquent.Table(e.TableName()).Where("menu_id = ? ", menu.MenuId).Update("paths", menu.Paths)
	return
}

func (e *MenuModel) Update(id int) (update MenuModel, err error) {
	if err = orm.Eloquent.Table(e.TableName()).First(&update, id).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Table(e.TableName()).Model(&update).Updates(&e).Error; err != nil {
		return
	}
	err = e.InitPaths(e)
	if err != nil {
		return
	}
	return
}

func (e *MenuModel) Delete(id int) (ok bool, err error) {
	if err = orm.Eloquent.Table(e.TableName()).Where("menu_id = ? ", id).Delete(&MenuModel{}).Error; err != nil {
		ok = false
		return
	}
	ok = true
	return
}

func (e *MenuModel) GetByRoleName(roleName string) (Menus []MenuModel, err error) {
	var roleMenModel = new(RoleMenuModel)
	roleMenuTable := roleMenModel.TableName()
	table := orm.Eloquent.Table(e.TableName()).Select(e.TableName() + ".*").Joins("left join " + roleMenuTable + " on " + roleMenuTable + ".menu_id = " + e.TableName() + ".menu_id")
	if err = table.Order("sort").Find(&Menus).Error; err != nil {
		return
	}

	return
}

func (m *MenuModel) GetPage() (Menus []MenuModel, err error) {
	table := orm.Eloquent.Table(m.TableName())
	if m.MenuName != "" {
		table = table.Where("name = ? ", m.MenuName)
	}
	if m.Title != "" {
		table = table.Where("title like ?", "%"+m.Title+"%")
	}
	if m.Visible != "" {
		table = table.Where("visible = ? ", m.Visible)
	}
	if m.MenuType != "" {
		table = table.Where("menu_type = ? ", m.MenuType)
	}

	if err = table.Order("sort").Find(&Menus).Error; err != nil {
		return
	}
	return
}
