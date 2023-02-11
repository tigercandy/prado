package monitor

import (
	"github.com/gin-gonic/gin"
	"github.com/tigercandy/prado/api/monitor"
	"github.com/tigercandy/prado/pkg/jwt"
)

func InitRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.RouterGroup {
	g := r.Group("")

	v1 := r.Group("/api/v1")
	v1.GET("/monitor/server", monitor.ServerInfo)

	return g
}
