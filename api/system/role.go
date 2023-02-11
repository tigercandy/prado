package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/tigercandy/prado/internal/models/common/response"
	"github.com/tigercandy/prado/internal/models/system"
	ss "github.com/tigercandy/prado/internal/services/system"
	"github.com/tigercandy/prado/pkg/utils"
)

func GetRoleList(c *gin.Context) {
	var (
		err      error
		page     = 1
		pageSize = 10
		data     system.SysRole
	)

	if offset := c.Request.FormValue("page"); offset != "" {
		page, _ = utils.String2Int(offset)
	}

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize, _ = utils.String2Int(size)
	}

	data.RoleKey = c.Request.FormValue("roleKey")
	data.RoleName = c.Request.FormValue("roleName")
	data.Status = c.Request.FormValue("status")

	result, total, err := data.GetPage(page, pageSize)
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	pageData := response.Page{
		Page:     page,
		PageSize: pageSize,
		Total:    total,
		List:     result,
	}
	response.PageResponse(c, pageData, "")
}

func GetRole(c *gin.Context) {
	var (
		err  error
		Role system.SysRole
	)
	Role.RoleId, _ = utils.String2Int(c.Param("roleId"))
	result, err := Role.Get()
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	menuIds, err := Role.GetRoleMenuId()
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	result.MenuIds = menuIds
	response.Success(c, result, "")
}

func CreateRole(c *gin.Context) {
	var data system.SysRole
	data.CreatedUser = ss.SysUserService.GetUserIdStr(c)
	err := c.BindWith(&data, binding.JSON)
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	id, err := data.Insert()
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	data.RoleId = id
	var t system.RoleMenuModel
	_, err = t.Insert(id, data.MenuIds)
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	response.Success(c, data, "")
}

func UpdateRole(c *gin.Context) {
	var (
		rm  system.RoleMenuModel
		r   system.SysRole
		err error
	)
	r.UpdatedUser = ss.SysUserService.GetUserIdStr(c)
	err = c.Bind(&r)
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	result, err := r.Update(r.RoleId)
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	_, err = rm.DeleteRoleMenu(r.RoleId)
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	_, err = rm.Insert(r.RoleId, r.MenuIds)
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	response.Success(c, result, "")
}

func DeleteRole(c *gin.Context) {
	var r system.SysRole
	r.UpdatedUser = ss.SysUserService.GetUserIdStr(c)

	ids := utils.IdsStr2IdsInt(c.Param("roleId"))
	_, err := r.BatchDelete(ids)
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	var rm system.RoleMenuModel
	_, err = rm.BatchDeleteRoleMenu(ids)
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	response.Success(c, "", "")
}
