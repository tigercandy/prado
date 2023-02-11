package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tigercandy/prado/internal/pkg/casbin"
	"github.com/tigercandy/prado/pkg/jwt"
	"net/http"
)

func AuthRoleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := c.Get("JWT_PAYLOAD")
		v := data.(jwt.MapClaims)
		e, _ := casbin.Casbin()
		res, _ := e.Enforce(v["roleKey"], c.Request.URL.Path, c.Request.Method)
		if res {
			c.Next()
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusForbidden,
				"msg":  "对不起，您没有访问权限，请联系管理员",
			})
			c.Abort()
			return
		}
	}
}
