package dao

import (
	"bluebull/model"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jmoiron/sqlx"
)

//这个是做注册时的数据操作
var db *sqlx.DB

//这个是在登陆后管理者对数据库的操作
var db1 *sqlx.DB

func Mysql(mysqlconf *model.MysqlConfig, department string) (err error) {
	//数据库的连接sqlx
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlconf.User,
		mysqlconf.Password,
		mysqlconf.Host,
		mysqlconf.Port,
		department,
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return model.ErrorMysqlConnect
	}
	//确定最大的容纳量和空闲量
	db.SetMaxOpenConns(mysqlconf.MaxOpen)
	db.SetMaxIdleConns(mysqlconf.MaxFree)

	return
}

func MysqlInit(mysqlconf *model.MysqlConfig) (err error) {
	//数据库的连接sqlx
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/register?charset=utf8&parseTime=True&loc=Local",
		mysqlconf.User,
		mysqlconf.Password,
		mysqlconf.Host,
		mysqlconf.Port,
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return model.ErrorMysqlConnect
	}
	//确定最大的容纳量和空闲量
	db.SetMaxOpenConns(mysqlconf.MaxOpen)
	db.SetMaxIdleConns(mysqlconf.MaxFree)

	return
}

func DBCloce() {
	db.Close()
}

func DB1Cloce() {
	db1.Close()
}
