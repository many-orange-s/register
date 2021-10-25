package model

import "github.com/pkg/errors"

//setting error
var (
	ErrorConfRead      = errors.New("读取文件配置错误")
	ErrorUnmarshalConf = errors.New("Conf反序列错误")
)

//Mysql error
var (
	ErrorMysqlConnect = errors.New("MySQL连接失败")
)

//logger error
var (
	ErrorUnmarshalLevel = errors.New("Logger反序列化错误")
)
