package model

import "github.com/pkg/errors"

//setting error
var (
	ErrorConfRead      = errors.New("读取文件配置错误")
	ErrorUnmarshalConf = errors.New("Conf反序列错误")
)

var ErrorMysqlConnect = errors.New("MySQL连接失败")

var ErrorUnmarshalLevel = errors.New("Logger反序列化错误")

//controller error
var (
	ErrorAdminExit = errors.New("用户已经存在")
	ErrorToken     = errors.New("token 鉴别失败")
)

var (
	ErrorRepeatModify = errors.New("重复修改")
	ErrorAdminNotExit = errors.New("用户不存在")
)
