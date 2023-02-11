package system

import (
	"github.com/gin-gonic/gin"
	"github.com/tigercandy/prado/internal/models/common/response"
	"github.com/tigercandy/prado/internal/models/system"
	ss "github.com/tigercandy/prado/internal/services/system"
)

func GetInfo(c *gin.Context) {
	var roles = make([]string, 1)
	roles[0] = ss.SysUserService.GetRoleName(c)

	var permissions = make([]string, 1)
	permissions[0] = "*:*:*"

	var buttons = make([]string, 1)
	buttons[0] = "*:*:*"

	RoleMenu := system.RoleMenuModel{}
	RoleMenu.RoleId = ss.SysUserService.GetRoleId(c)

	var mp = make(map[string]interface{})
	mp["roles"] = roles
	if ss.SysUserService.GetRoleName(c) == "admin" || ss.SysUserService.GetRoleName(c) == "系统管理员" {
		mp["permissions"] = permissions
		mp["buttons"] = buttons
	} else {
		list, _ := RoleMenu.GetPermissions()
		mp["permissions"] = list
		mp["buttons"] = list
	}

	sysuser := system.SysUser{}
	sysuser.UserId = ss.SysUserService.GetUserId(c)
	user, err := sysuser.Get()
	if err != nil {
		response.Error(c, -1, "未找到相关信息")
		return
	}
	mp["avatar"] = "https://s1.ax1x.com/2022/10/27/xfwt78.png"
	if user.Avatar != "" {
		mp["avatar"] = user.Avatar
	}
	mp["userName"] = user.NickName
	mp["userId"] = user.UserId
	mp["deptId"] = user.DeptId
	mp["name"] = user.NickName
	mp["introduction"] = "super administrator"

	response.Success(c, mp, "")
}
