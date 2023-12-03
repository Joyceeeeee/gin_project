package routers

import (
	"github.com/gin-gonic/gin"
	admin_v1 "youdangzhe/internal/controller/admin_v1"
	"youdangzhe/internal/middleware"
)

func SetAdminApiRoute(e *gin.Engine) {
	// version 1
	v1 := e.Group("api/v1")
	{
		// 无需校验权限
		loginC := admin_v1.NewLoginController()
		v1.POST("/login", loginC.Login)

		// 用户
		adminUser := v1.Group("users")
		{
			r := admin_v1.NewAdminUserController()
			// 获取用户信息
			adminUser.GET("", r.GetUserInfo)

			// 创建用户
			adminUser.POST("", r.RegisterUser)

			// 发送邮箱验证码
			adminUser.GET("/verify_code", admin_v1.NewEmailVerificationController().SendCode)
		}

		// 需要校验权限
		reqAuth := v1.Group("", middleware.AdminAuthHandler())
		{
			// api权限管理
			permissions := reqAuth.Group("permission")
			{
				r := admin_v1.NewPermissionController()
				permissions.POST("edit", r.Edit)
				permissions.GET("list", r.List)
			}

			// 菜单管理
			menu := reqAuth.Group("menu")
			{
				r := admin_v1.NewAdminUserController()
				menu.GET("get", r.GetUserInfo)
			}

			// 角色管理
			role := reqAuth.Group("role")
			{
				r := admin_v1.NewAdminUserController()
				role.GET("get", r.GetUserInfo)
			}

			// 用户组管理
			group := reqAuth.Group("group")
			{
				r := admin_v1.NewAdminUserController()
				group.GET("get", r.GetUserInfo)
			}
		}
	}
}
