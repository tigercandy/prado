package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/tigercandy/prado/internal/models/common/response"
	"github.com/tigercandy/prado/internal/models/system"
	ss "github.com/tigercandy/prado/internal/services/system"
	"github.com/tigercandy/prado/pkg/utils"
	"net/http"
)

func GetMenuRole(c *gin.Context) {
	result, err := ss.SysMenuService.SetMenuRole(ss.SysUserService.GetRoleName(c))
	if err != nil {
		response.Error(c, -1, "未找到相关信息")
		return
	}
	response.Success(c, result, "")
}

func GetMenuTreeRoleSelect(c *gin.Context) {
	var SysRole system.SysRole
	id, _ := utils.String2Int(c.Param("roleId"))
	SysRole.RoleId = id
	result, err := ss.SysMenuService.SetMenuLabel()
	if err != nil {
		response.Error(c, -1, "未找到相关信息")
		return
	}
	menuIds := make([]int, 0)
	if id != 0 {
		menuIds, err = SysRole.GetRoleMenuId()
		if err != nil {
			response.Error(c, -1, "未找到相关信息")
			return
		}
	}

	response.Custom(c, gin.H{
		"code":        http.StatusOK,
		"menus":       result,
		"checkedKeys": menuIds,
	})
}

func GetMenuIds(c *gin.Context) {
	var data system.RoleMenuModel
	data.RoleName = c.GetString("role")
	data.UpdatedUser = ss.SysUserService.GetUserIdStr(c)
	result, err := data.GetIds()
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	response.Success(c, result, "")
}

func GetMenu(c *gin.Context) {
	var data system.MenuModel
	id, _ := utils.String2Int(c.Param("id"))
	data.MenuId = id
	result, err := data.GetMenuById()
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	response.Success(c, result, "")
}

func MenuTreeSelect(c *gin.Context) {
	result, err := ss.SysMenuService.SetMenuLabel()
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	response.Success(c, result, "")
}

func CreateMenu(c *gin.Context) {
	var data system.MenuModel
	err := c.BindWith(&data, binding.JSON)
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	data.CreatedUser = ss.SysUserService.GetUserIdStr(c)
	result, err := data.Create()
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	response.Success(c, result, "")
}

func UpdateMenu(c *gin.Context) {
	var data system.MenuModel
	err := c.BindWith(&data, binding.JSON)
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	data.UpdatedUser = ss.SysUserService.GetUserIdStr(c)
	_, err = data.Update(data.MenuId)
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	response.Success(c, "", "")
}

func DeleteMenu(c *gin.Context) {
	var data system.MenuModel
	id, _ := utils.String2Int(c.Param("id"))

	data.UpdatedUser = ss.SysUserService.GetUserIdStr(c)
	_, err := data.Delete(id)
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	response.Success(c, "", "")
}

func GetMenuList(c *gin.Context) {
	var (
		err    error
		Menu   system.MenuModel
		result []system.MenuModel
	)
	Menu.MenuName = c.Request.FormValue("menuName")
	Menu.Visible = c.Request.FormValue("visible")
	Menu.Title = c.Request.FormValue("title")

	if Menu.Title == "" {
		result, err = ss.SysMenuService.SetMenu()
	} else {
		result, err = Menu.GetPage()
	}

	if err != nil {
		response.Error(c, -1, "")
		return
	}
	response.Success(c, result, "")
}
