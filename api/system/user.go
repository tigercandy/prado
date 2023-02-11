package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/tigercandy/prado/internal/models/common/response"
	"github.com/tigercandy/prado/internal/models/system"
	ss "github.com/tigercandy/prado/internal/services/system"
	"github.com/tigercandy/prado/pkg/utils"
	"net/http"
)

func GetUserList(c *gin.Context) {
	var (
		err      error
		page     = 1
		pageSize = 10
		data     system.SysUser
	)

	offset := c.Request.FormValue("page")
	if offset != "" {
		page, _ = utils.String2Int(offset)
	}

	size := c.Request.FormValue("pageSize")
	if size != "" {
		pageSize, _ = utils.String2Int(size)
	}

	data.Username = c.Request.FormValue("userName")
	data.NickName = c.Request.FormValue("nickName")
	data.Status = c.Request.FormValue("status")
	data.Phone = c.Request.FormValue("phone")

	postId := c.Request.FormValue("postId")
	data.PostId, _ = utils.String2Int(postId)

	deptId := c.Request.FormValue("deptId")
	data.DeptId, _ = utils.String2Int(deptId)

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

func GetUser(c *gin.Context) {
	var user system.SysUser
	user.UserId, _ = utils.String2Int(c.Param("userId"))
	result, err := user.Get()
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	var (
		role system.SysRole
		post system.PostModel
	)
	roles, _ := role.GetList()
	posts, _ := post.GetList()

	postIds := make([]int, 0)
	postIds = append(postIds, result.PostId)

	roleIds := make([]int, 0)
	roleIds = append(roleIds, result.RoleId)

	response.Custom(c, gin.H{
		"code":    http.StatusOK,
		"data":    result,
		"roleIds": roleIds,
		"postIds": postIds,
		"roles":   roles,
		"posts":   posts,
	})
}

func GetUserRolePost(c *gin.Context) {
	var (
		role system.SysRole
		post system.PostModel
	)
	roles, err := role.GetList()
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	posts, err := post.GetList()
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	mp := make(map[string]interface{}, 2)
	mp["roles"] = roles
	mp["posts"] = posts
	response.Success(c, mp, "")
}

func CreateUser(c *gin.Context) {
	var user system.SysUser
	err := c.MustBindWith(&user, binding.JSON)
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	user.CreatedUser = ss.SysUserService.GetUserIdStr(c)
	id, err := user.Insert()
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	response.Success(c, id, "")
}

func UpdateUser(c *gin.Context) {
	var user system.SysUser
	err := c.Bind(&user)
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	user.UpdatedUser = ss.SysUserService.GetUserIdStr(c)
	result, err := user.Update(user.UserId)
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	response.Success(c, result, "")
}

func DeleteUser(c *gin.Context) {
	var user system.SysUser
	user.UpdatedUser = ss.SysUserService.GetUserIdStr(c)
	ids := utils.IdsStr2IdsInt(c.Param("userId"))
	if result, err := user.BatchDelete(ids); err != nil {
		response.Error(c, -1, "")
		return
	} else {
		response.Success(c, result, "")
	}
}

func GetUserProfile(c *gin.Context) {
	var (
		dept system.DeptModel
		post system.PostModel
		role system.SysRole
		user system.SysUser
	)
	userId := ss.SysUserService.GetUserIdStr(c)
	user.UserId, _ = utils.String2Int(userId)
	result, err := user.Get()
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	roles, err := role.GetList()
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	posts, err := post.GetList()
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	dept.DeptId = result.DeptId
	d, err := dept.Get()
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	postIds := make([]int, 0)
	roleIds := make([]int, 0)

	postIds = append(postIds, result.PostId)
	roleIds = append(roleIds, result.RoleId)

	response.Custom(c, gin.H{
		"code":    http.StatusOK,
		"data":    result,
		"postIds": postIds,
		"roleIds": roleIds,
		"roles":   roles,
		"posts":   posts,
		"dept":    d,
	})
}

func UpdateUserAvatar(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	fileType := c.DefaultQuery("file_type", "images")
	if fileType != "images" {
		response.Error(c, -1, "文件格式错误!")
		return
	}
	files := form.File["upload[]"]
	guid := uuid.New().String()
	filePath := "static/upload/" + fileType + "/" + guid + ".jpg"
	for _, file := range files {
		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			response.Error(c, -1, "")
			return
		}
	}
	user := system.SysUser{}
	user.UserId = ss.SysUserService.GetUserId(c)
	user.Avatar = "/" + filePath
	user.UpdatedUser = ss.SysUserService.GetUserIdStr(c)
	_, _ = user.Update(user.UserId)
	response.Success(c, filePath, "")
}

func UpdateUserPwd(c *gin.Context) {

}
