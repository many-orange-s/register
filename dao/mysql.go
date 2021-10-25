package dao

import (
	"bluebull/model"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func MysqlInit(mysqlconf *model.MysqlConfig) (err error) {
	//数据库的连接sqlx
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlconf.User,
		mysqlconf.Password,
		mysqlconf.Host,
		mysqlconf.Port,
	//token
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

func Cloce() {
	db.Close()
}
