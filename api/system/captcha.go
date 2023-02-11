package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tigercandy/prado/internal/models/common/response"
	"github.com/tigercandy/prado/pkg/utils"
	"net/http"
)

func GenerateCaptchaHandler(c *gin.Context) {
	id, b64s, err := utils.DriverGigitFunc()
	if err != nil {
		response.Error(c, -1, fmt.Sprintf("验证码获取失败, %v", err.Error()))
		return
	}
	response.Custom(c, gin.H{
		"code": http.StatusOK,
		"data": b64s,
		"id":   id,
		"msg":  "ok",
	})
}
