package dao

import (
	"bluebull/model"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func ShowAllData(department string) (allMsg []*model.AllMsg, err error) {
	sqlStr := fmt.Sprintf("select * from %s", department)
	db.Rebind(sqlStr)
	err = db.Select(&allMsg, sqlStr)
	if err != nil {
		return nil, err
	}
	return
}

func ShowAData(department string, name string) (Msg []*model.AllMsg, err error) {
	sqlStr := fmt.Sprintf("select * from %s where name in (?)", department)
	db.Rebind(sqlStr)
	err = db.Select(&Msg, sqlStr, name)
	if err != nil {
		return nil, err
	}
	return Msg, err
}

func SearchGroup(department string, groupname string) (Msg []*model.AllMsg, err error) {
	sqlStr := fmt.Sprintf("select * from %s where group_name = ?", department)
	db.Rebind(sqlStr)
	err = db.Select(&Msg, sqlStr, groupname)
	if err != nil {
		return nil, err
	}
	return Msg, err
}

func Updata(department string, msg *model.Update, id string) (ret sql.Result, err error) {
	sqlStr := fmt.Sprintf("update %s set %s = ? where id = ?", department, msg.Target)
	db.Rebind(sqlStr)
	ret, err = db.Exec(sqlStr, msg.UpdateData, id)
	return
}

func AddData(department string, msg []interface{}) (err error) {
	sqlStr := fmt.Sprintf("insert into %s (name,gender,grade,birth,telephone,group_name) values (?)", department)
	query, args, err := sqlx.In(sqlStr, msg)
	if err != nil {
		return
	}

	db.Rebind(query)
	_, err = db.Exec(query, args...)
	return
}

func Delete(department string, id string) (ret sql.Result, err error) {
	sqlStr := fmt.Sprintf("delete from %s where id = ?", department)
	ret, err = db.Exec(sqlStr, id)
	return
}
