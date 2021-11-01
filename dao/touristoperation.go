package dao

import (
	"bluebull/model"
	"fmt"
	"io/ioutil"
)

// SelectAdmin 找到用户 返回表中的department
func SelectAdmin(msg *model.Sign) (department string, err error) {
	sqlStr := `select department from admin where name = ? and password = ?`
	err = db.Get(&department, sqlStr, msg.Name, msg.Password)
	return
}

// SearchSame 返回用户名数目
func SearchSame(name string) (exit bool, err error) {
	var count int
	sqlStr := `select Count(name) from admin where name = ?`
	if err = db.Get(&count, sqlStr, name); err != nil {
		return
	}

	if count > 0 {
		exit = false
	} else {
		exit = true
	}
	return
}

// SignUp 插入用户信息
func SignUp(reg *model.Register) (err error) {
	sqlStr := `insert into admin(name,password,department) values (?,?,?)`
	_, err = db.Exec(sqlStr, reg.Name, reg.Password, reg.Department)
	return
}

// CreateForm 建表
func CreateForm(department string) (err error) {
	sqlBytes, err := ioutil.ReadFile("./table.sql")
	if err != nil {
		return model.ErrorTableReadFail
	}
	sqlStr := string(sqlBytes)
	sqlStr = fmt.Sprintf(sqlStr, department)
	_, err = db.Exec(sqlStr)
	return
}
