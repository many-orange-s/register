package main

import (
	"bluebull/dao"
	"bluebull/logger"
	"bluebull/router"
	"bluebull/setting"
	"go.uber.org/zap"
	"log"
)

func main() {
	//配置文件的读取
	err := setting.Init()
	if err != nil {
		log.Printf("%+v", err)
		return
	}

	//日志设置
	err = logger.Init(setting.Conf.LogConfig, setting.Conf.Mode)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer zap.L().Sync()

	//mysql的初始化
	err = dao.MysqlInit(setting.Conf.MysqlConfig)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	defer dao.DBCloce()
	//redis的初始化

	//router的使用
	r := router.SetUpRouter()
	r.Run(":8080")
}
