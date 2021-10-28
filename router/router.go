package router

import (
	"bluebull/controller"
	"bluebull/logger"
	"bluebull/middleware"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecover(true))

	//管理者的登录和注册
	loginGroup := r.Group("/tourist")
	{
		//登录
		loginGroup.POST("/sign", controller.AdminSign)
		//注册
		loginGroup.POST("/register", controller.AddAdmin)
	}

	oprationGroup := r.Group("/admin")
	oprationGroup.Use(middleware.TokenAuthMiddle())
	{

	}
	return r
}
