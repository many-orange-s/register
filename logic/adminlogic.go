package logic

import (
	"bluebull/dao"
	"bluebull/model"
)

func Updata(department string, NewMsg *model.Update, id string) (err error) {
	aff, err := dao.Updata(department, NewMsg, id)
	if aff == 0 {
		return model.ErrorNotExit
	}
	return
}
