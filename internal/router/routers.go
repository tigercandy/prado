package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tigercandy/prado/internal/middleware"
	"github.com/tigercandy/prado/internal/router/dashboard"
	"github.com/tigercandy/prado/internal/router/monitor"
	"github.com/tigercandy/prado/internal/router/system"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	// 初始化中间件
	middleware.InitMiddleware(r)
	// 初始化权限
	authMiddleware, err := middleware.InitAuthMiddleware()
	if err != nil {
		panic(err)
	}
	// 首页路由
	dashboard.InitRouter(r, authMiddleware)
	// 系统路由
	system.InitRouter(r, authMiddleware)
	// 系统监控路由
	monitor.InitRouter(r, authMiddleware)

	return r
}
