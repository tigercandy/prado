package system

import (
	"github.com/gin-gonic/gin"
	fileApi "github.com/tigercandy/prado/api/file"
	"github.com/tigercandy/prado/api/system"
	"github.com/tigercandy/prado/internal/handler"
	"github.com/tigercandy/prado/internal/middleware"
	"github.com/tigercandy/prado/pkg/jwt"
)

func InitRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.RouterGroup {
	g := r.Group("")
	// 静态资源
	staticFileRouter(g, r)
	// 无权限认证路由
	unAuthRouter(g)
	// 权限认证路由
	authRouter(g, authMiddleware)

	return g
}

func staticFileRouter(r *gin.RouterGroup, g *gin.Engine) {
	r.Static("/static", "./static")
	g.LoadHTMLGlob("web/index.html")
}

func unAuthRouter(r *gin.RouterGroup) {
	v1 := r.Group("/api/v1")
	v1.GET("/settings", system.GetSettings)
	v1.GET("/getCaptcha", system.GenerateCaptchaHandler)
	v1.GET("/menuTreeSelect", system.MenuTreeSelect)

	registerFileRouter(v1)
}

func authRouter(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	r.POST("/login", authMiddleware.LoginHandler)
	r.GET("/refresh_token", authMiddleware.RefreshHandler)

	v1 := r.Group("/api/v1")
	// 系统路由
	registerBaseRouter(v1, authMiddleware)
	// 页面路由
	registerPageRouter(v1, authMiddleware)
	// 用户路由
	registerUserRouter(v1, authMiddleware)
	// 部门路由
	registerDeptRouter(v1, authMiddleware)
	// 菜单路由
	registerMenuRouter(v1, authMiddleware)
	// 角色路由
	registerRoleRouter(v1, authMiddleware)
	// 岗位路由
	registerPostRouter(v1, authMiddleware)
	// 系统设置
	registerSettingsRouter(v1, authMiddleware)
}

// 基础路由
func registerBaseRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	v1auth := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthRoleMiddleware())
	{
		v1auth.POST("/logout", handler.Logout)
		v1auth.GET("/getInfo", system.GetInfo)
		v1auth.GET("/menuRole", system.GetMenuRole)
		v1auth.GET("/menuIds", system.GetMenuIds)
		v1auth.GET("/roleMenuTreeSelect/:roleId", system.GetMenuTreeRoleSelect)
		v1auth.GET("/roleDeptTreeSelect/:roleId", system.GetDeptTreeRoleSelect)
	}
}

// 页面路由
func registerPageRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	v1auth := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthRoleMiddleware())
	{
		v1auth.GET("/deptList", system.GetDeptList)
		v1auth.GET("/ordinaryDeptList", system.GetOrdinaryDeptList)
		v1auth.GET("/deptTree", system.DeptTree)
		v1auth.GET("/userList", system.GetUserList)
		v1auth.GET("/roleList", system.GetRoleList)
		v1auth.GET("/postList", system.GetPostList)
		v1auth.GET("/menuList", system.GetMenuList)
	}
}

// 用户路由
func registerUserRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	user := v1.Group("/user").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthRoleMiddleware())
	{
		user.GET("/:userId", system.GetUser)
		user.GET("/", system.GetUserRolePost)
		user.POST("", system.CreateUser)
		user.PUT("", system.UpdateUser)
		user.DELETE("/:userId", system.DeleteUser)
		user.GET("/profile", system.GetUserProfile)
		user.POST("/avatar", system.UpdateUserAvatar)
		user.PUT("/pwd", system.UpdateUserPwd)
	}
}

// 部门路由
func registerDeptRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	dept := v1.Group("/dept").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthRoleMiddleware())
	{
		dept.GET("/:deptId", system.GetDept)
		dept.POST("", system.CreateDept)
		dept.PUT("", system.UpdateDept)
		dept.DELETE("/:deptId", system.DeleteDept)
	}
}

// 菜单路由
func registerMenuRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	menu := v1.Group("/menu").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthRoleMiddleware())
	{
		menu.GET("/:id", system.GetMenu)
		menu.POST("", system.CreateMenu)
		menu.PUT("", system.UpdateMenu)
		menu.DELETE("/:id", system.DeleteMenu)
	}
}

// 角色路由
func registerRoleRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	role := v1.Group("/role").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthRoleMiddleware())
	{
		role.GET("/:roleId", system.GetRole)
		role.POST("", system.CreateRole)
		role.PUT("", system.UpdateRole)
		role.DELETE("/:roleId", system.DeleteRole)
	}
}

// 岗位路由
func registerPostRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	post := v1.Group("/post").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthRoleMiddleware())
	{
		post.GET("/:postId", system.GetPost)
		post.POST("", system.CreatePost)
		post.PUT("", system.UpdatePost)
		post.DELETE("/:postId", system.DeletePost)
	}
}

// 系统设置路由
func registerSettingsRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	settings := v1.Group("/settings").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthRoleMiddleware())
	{
		settings.POST("", system.SetSettings)
	}
}

// 文件上传路由
func registerFileRouter(v1 *gin.RouterGroup) {
	file := v1.Group("/file")
	{
		file.POST("/upload", fileApi.UploadFile)
	}
}
