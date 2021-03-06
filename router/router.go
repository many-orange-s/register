package router

import (
	"bluebull/controller"
	"bluebull/logger"
	"bluebull/middleware"
	"github.com/gin-gonic/gin"
)

// SetUpRouter 路由
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

	operationGroup := r.Group("/admin")
	operationGroup.Use(middleware.TokenAuthMiddle())
	{
		//返回表中的所有信息
		operationGroup.GET("/show", controller.ShowAllData)
		//返回表中的一条信息
		operationGroup.GET("/search/:name", controller.ShowAData)
		//返回表中一个组的元素
		operationGroup.GET("/search/group/:groupname", controller.SearchGroup)
		//修改表中的一条数据
		operationGroup.PUT("/update/:id", controller.UpdateData)
		//添加表中的元素
		operationGroup.POST("/add", controller.Add)
		//删除一条元素
		operationGroup.DELETE("/delete/:id", controller.DeleteData)
	}
	return r
}
