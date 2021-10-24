package main

import (
	"bluebull/dao"
	"bluebull/logger"
	"bluebull/router"
	"bluebull/setting"
	"go.uber.org/zap"
)

func main() {
	//配置文件的读取
	err := setting.Init()
	if err != nil {

	}

	//日志设置
	err = logger.Init(setting.Conf.LogConfig)
	if err != nil {

	}
	defer zap.L().Sync()

	//mysql的初始化
	err = dao.MysqlInit(setting.Conf.MysqlConfig)
	if err != nil {

	}
	defer dao.Cloce()
	//redis的初始化

	//router的使用
	r := router.SetUpRouter()
	r.Run(":8080")
}
