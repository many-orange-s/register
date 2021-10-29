package logic

import (
	"bluebull/dao"
	"bluebull/model"
)

func Updata(department string, NewMsg *model.Update, id string) (err error) {
	ret, err := dao.Updata(department, NewMsg, id)
	//这个错误判断是看exec是否错误的
	if err != nil {
		return
	}
	//这个错误判断是看是否你的id修改是否存在 或 重复修改
	aff, err := ret.RowsAffected()
	if aff == 0 {
		return model.ErrorRepeatModify
	}
	return
}

func AddMember(department string, msg *model.AllMsg) (err error) {
	var inf []interface{}
	inf = append(inf, msg.Name, msg.Gender, msg.Grade, msg.Birth, msg.Telephone, msg.GroupName)
	err = dao.AddData(department, inf)
	return
}
