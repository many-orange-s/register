package dao

import (
	"bluebull/model"
	"database/sql"
	"fmt"
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
	sqlStr := fmt.Sprintf("select * from %s where name = ?", department)
	db.Rebind(sqlStr)
	err = db.Select(&Msg, sqlStr, name)
	if err != nil {
		return nil, err
	}
	return Msg, err
}

func SearchData(department string, id string) (Msg *model.AllMsg, err error) {
	sqlStr := fmt.Sprintf("select * from %s where id = ?", department)
	db.Rebind(sqlStr)
	err = db.Get(&Msg, sqlStr, id)
	if err != nil {
		return nil, err
	}
	return Msg, err
}

func Updata(department string, msg *model.Update, id string) (ret sql.Result, err error) {
	sqlStr := fmt.Sprintf("update %s set %s = ? where id = ?", department, msg.Target)
	db.Rebind(sqlStr)
	ret, err = db.Exec(sqlStr, msg.UpdateDate, id)
	return
}
