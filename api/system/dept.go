package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/tigercandy/prado/global/orm"
	"github.com/tigercandy/prado/internal/models/common/response"
	"github.com/tigercandy/prado/internal/models/system"
	ss "github.com/tigercandy/prado/internal/services/system"
	"github.com/tigercandy/prado/pkg/utils"
)

func GetDeptTreeRoleSelect(c *gin.Context) {

}

func GetDeptList(c *gin.Context) {
	var (
		Dept   system.DeptModel
		result []system.DeptModel
		err    error
	)
	Dept.DeptName = c.Request.FormValue("deptName")
	Dept.Status = c.Request.FormValue("status")
	Dept.DeptId, _ = utils.String2Int(c.Request.FormValue("deptId"))

	if Dept.DeptName == "" {
		result, err = ss.SysDeptService.SetDept()
	} else {
		result, err = Dept.GetPage()
	}
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	response.Success(c, result, "")
}

func GetOrdinaryDeptList(c *gin.Context) {
	var (
		err      error
		deptList []system.DeptModel
	)
	err = orm.Eloquent.Model(&system.DeptModel{}).Find(&deptList).Error
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	response.Success(c, deptList, "")
}

func DeptTree(c *gin.Context) {
	var (
		err  error
		Dept system.DeptModel
	)
	Dept.DeptName = c.Request.FormValue("deptName")
	Dept.Status = c.Request.FormValue("status")
	Dept.DeptId, _ = utils.String2Int(c.Request.FormValue("deptId"))

	result, err := ss.SysDeptService.SetDept()
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	response.Success(c, result, "")
}

func GetDept(c *gin.Context) {
	var (
		err  error
		dept system.DeptModel
	)
	dept.DeptId, _ = utils.String2Int(c.Param("deptId"))
	result, err := dept.Get()
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	response.Success(c, result, "")
}

func CreateDept(c *gin.Context) {
	var dept system.DeptModel
	err := c.BindWith(&dept, binding.JSON)
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	dept.CreatedUser = ss.SysUserService.GetUserIdStr(c)
	result, err := dept.Create()
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	response.Success(c, result, "")
}

func UpdateDept(c *gin.Context) {
	var dept system.DeptModel
	err := c.BindJSON(&dept)
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	dept.UpdatedUser = ss.SysUserService.GetUserIdStr(c)
	result, err := dept.Update(dept.DeptId)
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	response.Success(c, result, "")
}

func DeleteDept(c *gin.Context) {
	var dept system.DeptModel
	id, _ := utils.String2Int(c.Param("id"))
	_, err := dept.Delete(id)
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	response.Success(c, "", "")
}
