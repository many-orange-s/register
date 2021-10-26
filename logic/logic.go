package logic

import (
	"bluebull/dao"
	"bluebull/model"
)

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
