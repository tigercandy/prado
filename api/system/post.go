package system

import (
	"github.com/gin-gonic/gin"
	"github.com/tigercandy/prado/internal/models/common/response"
	"github.com/tigercandy/prado/internal/models/system"
	ss "github.com/tigercandy/prado/internal/services/system"
	"github.com/tigercandy/prado/pkg/utils"
)

func GetPostList(c *gin.Context) {
	var (
		err      error
		data     system.PostModel
		page     = 1
		pageSize = 10
	)

	if offset := c.Request.FormValue("page"); offset != "" {
		page, _ = utils.String2Int(offset)
	}

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize, _ = utils.String2Int(size)
	}

	data.PostId, _ = utils.String2Int(c.Request.FormValue("postId"))
	data.PostCode = c.Request.FormValue("postCode")
	data.PostName = c.Request.FormValue("postName")
	data.Status = c.Request.FormValue("status")

	result, total, err := data.GetPage(page, pageSize)
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	pageData := response.Page{
		page,
		pageSize,
		total,
		result,
	}
	response.PageResponse(c, pageData, "")
}

func GetPost(c *gin.Context) {
	var (
		err  error
		Post system.PostModel
	)
	Post.PostId, err = utils.String2Int(c.Param("postId"))

	result, err := Post.Get()
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	response.Success(c, result, "")
}

func CreatePost(c *gin.Context) {
	var data system.PostModel
	err := c.Bind(&data)
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	data.CreatedUser = ss.SysUserService.GetUserIdStr(c)
	data.CreatedAt = utils.GetCurrentTime()
	result, err := data.Create()
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	response.Success(c, result, "")
}
func UpdatePost(c *gin.Context) {
	var data system.PostModel
	err := c.Bind(&data)
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	data.UpdatedUser = ss.SysUserService.GetUserIdStr(c)
	data.UpdatedAt = utils.GetCurrentTime()
	result, err := data.Update(data.PostId)
	if err != nil {
		response.Error(c, -1, "")
		return
	}
	response.Success(c, result, "")
}

func DeletePost(c *gin.Context) {
	var data system.PostModel
	data.UpdatedUser = ss.SysUserService.GetUserIdStr(c)
	data.UpdatedAt = utils.GetCurrentTime()
	ids := utils.IdsStr2IdsInt(c.Param("postIds"))
	if result, err := data.BatchDelete(ids); err != nil {
		response.Error(c, -1, "")
		return
	} else {
		response.Success(c, result, "")
	}
}
