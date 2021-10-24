package router

import (
	"bluebull/logger"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecover(true))

	r.GET()
	return r
}
