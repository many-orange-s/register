package dao

import (
	"bluebull/model"
)

func SelectAdmin(msg *model.Sign) (department string, err error) {
	sqlStr := `select department from admin where name = ? and password = ?`
	err = db.Get(&department, sqlStr, msg.Name, msg.Password)
	return
}

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

func SignUp(reg *model.Register) (err error) {
	sqlStr := `insert into admin(name,password,department) values (?,?,?)`
	_, err = db.Exec(sqlStr, reg.Name, reg.Password, reg.Department)
	return
}
