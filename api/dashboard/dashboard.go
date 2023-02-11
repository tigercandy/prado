package dashboard

import (
	"github.com/gin-gonic/gin"
	"github.com/tigercandy/prado/internal/models/common/response"
	"github.com/tigercandy/prado/internal/services/dashboard"
)

func InitData(c *gin.Context) {
	var (
		count  map[string]int
		ranks  []dashboard.Ranks
		submit map[string][]interface{}
		handle interface{}
		period interface{}
	)
	response.Success(c, map[string]interface{}{
		"count":  count,
		"ranks":  ranks,
		"submit": submit,
		"handle": handle,
		"period": period,
	}, "")
}
