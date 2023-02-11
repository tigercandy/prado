package system

import (
	"fmt"
	"github.com/tigercandy/prado/global/orm"
	"github.com/tigercandy/prado/pkg/utils"
)

type RoleMenuModel struct {
	Id          int    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id" form:"id"`
	RoleId      int    `gorm:"type:int(11)"`
	MenuId      int    `gorm:"type:int(11)"`
	RoleName    string `gorm:"type:varchar(128)"`
	CreatedUser string `gorm:"type:varchar(128)"`
	UpdatedUser string `gorm:"type:varchar(128)"`
}

type MenuPath struct {
	Path string `json:"path"`
}

func (RoleMenuModel) TableName() string {
	return "sys_role_menu"
}

func (rm *RoleMenuModel) Get() ([]RoleMenuModel, error) {
	var r []RoleMenuModel
	table := orm.Eloquent.Table(rm.TableName())
	if rm.RoleId != 0 {
		table = table.Where("role_id = ? ", rm.RoleId)
	}

	if err := table.Find(&r).Error; err != nil {
		return nil, err
	}

	return r, nil
}

func (rm *RoleMenuModel) GetPermissions() ([]string, error) {
	var r []MenuModel
	var menuModel = new(MenuModel)
	menuTable := menuModel.TableName()
	table := orm.Eloquent.Select(menuTable + ".permission").Table(menuTable).Joins("left join " + rm.TableName() + " on " + menuTable + ".menu_id = " + rm.TableName() + ".menu_id")
	table = table.Where("role_id = ? ", rm.RoleId)
	table = table.Where(menuTable + ".menu_type in('F', 'C')")
	if err := table.Find(&r).Error; err != nil {
		return nil, err
	}
	var list []string
	for i := 0; i < len(r); i++ {
		list = append(list, r[i].Permission)
	}
	return list, nil
}

func (rm *RoleMenuModel) GetIds() ([]MenuPath, error) {
	var r []MenuPath
	var menuModel = new(MenuModel)
	var roleModel = new(SysRole)
	menuTable := menuModel.TableName()
	roleTable := roleModel.TableName()
	table := orm.Eloquent.Select(menuTable + ".path").Table(rm.TableName())
	table = table.Joins("left join " + roleTable + " on " + roleTable + ".role_id = " + rm.TableName() + ".role_id")
	table = table.Joins("left join " + menuTable + " on " + menuTable + ".menu_id = " + rm.TableName() + ".menu_id")
	table = table.Where(roleTable+".role_name = ? and "+menuTable+".type=1", rm.RoleName)
	if err := table.Find(&r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func (rm *RoleMenuModel) Insert(roleId int, menuId []int) (bool, error) {
	var role SysRole
	if err := orm.Eloquent.Table(role.TableName()).Where("role_id = ? ", roleId).First(&role).Error; err != nil {
		return false, err
	}
	var (
		menu      []MenuModel
		menuModel MenuModel
	)
	if err := orm.Eloquent.Table(menuModel.TableName()).Where("menu_id in (?) ", menuId).Find(&menu).Error; err != nil {
		return false, err
	}
	var casbinModel CasbinRule
	sql1 := "INSERT INTO " + rm.TableName() + " (`role_id`, `menu_id`, `role_name`) VALUES "
	sql2 := "INSERT INTO " + casbinModel.TableName() + " (`p_type`, `v0`, `v1`, `v2`) VALUES "

	for i := 0; i < len(menu); i++ {
		if len(menu)-1 == i {
			sql1 += fmt.Sprintf("(%d%d, '%s');", role.RoleId, menu[i], role.RoleKey)
			if menu[i].MenuType == "A" {
				sql2 += fmt.Sprintf("('p', '%s', '%s', '%s');", role.RoleKey, menu[i].Path, menu[i], menu[i].Action)
			}
		} else {
			sql1 += fmt.Sprintf("(%d,%d, '%s'),", role.RoleId, menu[i].MenuId, role.RoleKey)
			if menu[i].MenuType == "A" {
				sql2 += fmt.Sprintf("('p', '%s', '%s', '%s'),", role.RoleKey, menu[i].Path, menu[i], menu[i].Action)
			}
		}
	}
	orm.Eloquent.Exec(sql1)
	sql2 = sql2[0:len(sql2)-1] + ";"
	orm.Eloquent.Exec(sql2)

	return true, nil
}

func (rm *RoleMenuModel) Delete(roleId, menuId string) (bool, error) {
	rm.RoleId, _ = utils.String2Int(roleId)
	table := orm.Eloquent.Table(rm.TableName()).Where("role_id = ? ", roleId)
	if menuId != "" {
		table = table.Where("menu_id = ? ", menuId)
	}
	if err := table.Delete(&rm).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (rm *RoleMenuModel) BatchDeleteRoleMenu(roleIds []int) (bool, error) {
	if err := orm.Eloquent.Table(rm.TableName()).Where("role_id in (?) ", roleIds).Delete(&rm).Error; err != nil {
		return false, err
	}
	var (
		role      []SysRole
		roleModel SysRole
	)
	if err := orm.Eloquent.Table(roleModel.TableName()).Where("role_id in (?) ", roleIds).Find(&role).Error; err != nil {
		return false, err
	}
	sql := ""
	var casbinModel CasbinRule
	for i := 0; i < len(role); i++ {
		sql += "DELETE FROM " + casbinModel.TableName() + " WHERE v0 = '" + role[i].RoleName + "';"
	}
	orm.Eloquent.Exec(sql)
	return true, nil
}

func (rm *RoleMenuModel) DeleteRoleMenu(roleId int) (bool, error) {
	var roleDept RoleDeptMode
	if err := orm.Eloquent.Table(roleDept.TableName()).Where("role_id = ? ", roleId).Delete(&rm).Error; err != nil {
		return false, err
	}
	if err := orm.Eloquent.Table(rm.TableName()).Where("role_id = ? ", roleId).Delete(&rm).Error; err != nil {
		return false, err
	}
	var role SysRole
	if err := orm.Eloquent.Table(role.TableName()).Where("role_id = ? ", roleId).First(&role).Error; err != nil {
		return false, err
	}
	var casbinModel CasbinRule
	sql := "DELETE FROM " + casbinModel.TableName() + " WHERE v0 = '" + role.RoleKey + "';"
	orm.Eloquent.Exec(sql)
	return true, nil
}
