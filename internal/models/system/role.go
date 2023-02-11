package system

import (
	"errors"
	"github.com/tigercandy/prado/global/orm"
	"github.com/tigercandy/prado/pkg/utils"
)

type RoleModel struct {
	RoleId      int    `gorm:"primary_key;AUTO_INCREMENT" json:"roleId"`
	RoleName    string `gorm:"type:varchar(128)" json:"roleName"`
	Status      string `gorm:"type:int(1)" json:"status"`
	RoleKey     string `gorm:"type:varchar(128)" json:"roleKey"`
	RoleSort    int    `gorm:"type:int(4)" json:"roleSort"`
	Flag        string `gorm:"type:varchar(128)" json:"flag"`
	CreatedUser string `gorm:"type:varchar(128)" json:"createdUser"`
	UpdatedUser string `gorm:"type:varchar(128)" json:"updatedUser"`
	Remark      string `gorm:"type:varchar(255)" json:"remark"`
	Admin       bool   `gorm:"type:char(1)" json:"admin"`
	MenuIds     []int  `gorm:"-" json:"menuIds"`
	DeptIds     []int  `gorm:"-" json:"deptIds"`
	Params      string `gorm:"-" json:"params"`
	BaseModel
}

type SysRole struct {
	RoleModel
}

type MenuIdList struct {
	MenuId int `json:"menuId"`
}

func (SysRole) TableName() string {
	return "sys_role"
}

func (r *SysRole) GetRoleMenuId() ([]int, error) {
	var roleMenuModel = new(RoleMenuModel)
	var menuModel = new(MenuModel)
	roleMenuTable := roleMenuModel.TableName()
	menuTable := menuModel.TableName()
	menuIds := make([]int, 0)
	menuList := make([]MenuIdList, 0)
	if err := orm.Eloquent.Table(roleMenuTable).Select(roleMenuTable+".menu_id").Joins("left join "+menuTable+" on "+menuTable+".menu_id = "+roleMenuTable+".menu_id").Where("role_id = ? ", r.RoleId).Where(roleMenuTable+".menu_id not in (select "+menuTable+".parent_id from "+roleMenuTable+" left join "+menuTable+" on "+menuTable+".menu_id="+roleMenuTable+".menu_id where role_id = ? )", r.RoleId).Find(&menuList).Error; err != nil {
		return nil, err
	}

	for i := 0; i < len(menuList); i++ {
		menuIds = append(menuIds, menuList[i].MenuId)
	}

	return menuIds, nil
}

func (r *SysRole) GetPage(page, pageSize int) ([]SysRole, int, error) {
	var (
		doc   []SysRole
		total int64
	)

	table := orm.Eloquent.Select("*").Table(r.TableName())
	if r.RoleId != 0 {
		table = table.Where("role_id = ? ", r.RoleId)
	}
	if r.RoleName != "" {
		table = table.Where("role_name like ? ", "%"+r.RoleName+"%")
	}
	if r.Status != "" {
		table = table.Where("status = ? ", r.Status)
	}
	if r.RoleKey != "" {
		table = table.Where("role_key like ? ", "%"+r.RoleKey+"%")
	}

	if err := table.Order("role_sort").Offset((page - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("`deleted_at` IS NULL").Count(&total)
	return doc, utils.Int642Int(total), nil
}

func (r *SysRole) Get() (SysRole SysRole, err error) {
	table := orm.Eloquent.Table(r.TableName())
	if r.RoleId != 0 {
		table = table.Where("role_id = ? ", r.RoleId)
	}
	if r.RoleName != "" {
		table = table.Where("role_name = ? ", r.RoleName)
	}
	if err = table.First(&SysRole).Error; err != nil {
		return
	}
	return
}

func (r *SysRole) GetList() (role []SysRole, err error) {
	table := orm.Eloquent.Table(r.TableName())
	if r.RoleId != 0 {
		table = table.Where("role_id = ? ", r.RoleId)
	}
	if r.RoleName != "" {
		table = table.Where("role_name = ? ", r.RoleName)
	}
	if err = table.Order("role_sort").First(&role).Error; err != nil {
		return
	}
	return
}

func (r *SysRole) Insert() (id int, err error) {
	var count int64
	orm.Eloquent.Table(r.TableName()).Where("role_name = ? or role_key = ? ) and `deleted_at` IS NULL", r.RoleName, r.RoleKey).Count(&count)
	if count > 0 {
		return 0, errors.New("角色已存在，无法创建!")
	}
	r.UpdatedUser = ""
	if err = orm.Eloquent.Table(r.TableName()).Create(&r).Error; err != nil {
		return
	}
	id = r.RoleId
	return
}

func (r *SysRole) Update(id int) (update SysRole, err error) {
	if err = orm.Eloquent.Table(r.TableName()).First(&update, id).Error; err != nil {
		return
	}
	if r.RoleName != "" && r.RoleKey != update.RoleKey {
		return update, errors.New("角色标识无法修改!")
	}
	if err = orm.Eloquent.Table(r.TableName()).Model(&update).Updates(&r).Error; err != nil {
		return
	}
	return
}

func (r *SysRole) BatchDelete(id []int) (res bool, err error) {
	if err = orm.Eloquent.Table(r.TableName()).Where("role_id in (?)", id).Delete(&SysRole{}).Error; err != nil {
		return
	}
	res = true
	return
}
