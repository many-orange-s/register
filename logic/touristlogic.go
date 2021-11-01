package logic

import (
	"bluebull/dao"
	"bluebull/model"
)

// Logon 判断用户是否已存在
func Logon(reg *model.Register) error {
	exit, err := dao.SearchSame(reg.Name)
	if err != nil {
		return err
	}

	if exit == false {
		return model.ErrorAdminExit
	} else {
		return nil
	}
}
