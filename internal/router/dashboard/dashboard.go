package dashboard

import (
	"github.com/gin-gonic/gin"
	"github.com/tigercandy/prado/api/dashboard"
	"github.com/tigercandy/prado/internal/middleware"
	"github.com/tigercandy/prado/pkg/jwt"
)

func InitRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.RouterGroup {
	g := r.Group("")
	v1 := r.Group("/api/v1")
	classify := v1.Group("/dashboard").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthRoleMiddleware())
	{
		classify.GET("", dashboard.InitData)
	}

	return g
}
