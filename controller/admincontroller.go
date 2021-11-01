package controller

import (
	"bluebull/dao"
	"bluebull/logic"
	"bluebull/model"
	"bluebull/respond"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ShowAllData 展示表中所有信息
func ShowAllData(c *gin.Context) {
	department := c.GetString("department")
	allMsg, err := dao.ShowAllData(department)
	if err != nil {
		zap.L().Error("dao.ShowAllDate err", zap.Error(err))
		respond.Fail(c, respond.CodeSystemBusy)
		return
	}
	respond.SuccessWith(c, allMsg)
}

// ShowAData 展示一种信息
func ShowAData(c *gin.Context) {
	name, ok := c.Params.Get("name")
	if !ok {
		respond.Fail(c, respond.CodeMsgNotReceive)
		return
	}

	department := c.GetString("department")
	Msg, err := dao.ShowAData(department, name)
	if err != nil {
		zap.L().Error("dao.ShowAData err", zap.Error(err))
		respond.Fail(c, respond.CodeSystemBusy)
		return
	}
	if Msg == nil && err == nil {
		respond.Fail(c, respond.CodeHasNotExit)
		return
	}
	respond.SuccessWith(c, Msg)
}

// UpdateData 更新信息
func UpdateData(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		respond.Fail(c, respond.CodeMsgNotReceive)
		return
	}

	department := c.GetString("department")

	NewMsg := new(model.Update)
	err := c.BindJSON(NewMsg)
	if err != nil {
		respond.Fail(c, respond.CodeParamInvalid)
		return
	}

	err = logic.Updata(department, NewMsg, id)
	if err != nil {
		if errors.Is(err, model.ErrorRepeatModify) {
			respond.Fail(c, respond.CodeRepeatModify)
			return
		}
		zap.L().Error("dao.Update err", zap.Error(err))
		respond.Fail(c, respond.CodeSystemBusy)
		return
	}
	respond.Success(c)
}

// SearchGroup 用小组来展示信息
func SearchGroup(c *gin.Context) {
	groupname, ok := c.Params.Get("groupname")
	if !ok {
		respond.Fail(c, respond.CodeMsgNotReceive)
		return
	}

	department := c.GetString("department")
	msg, err := dao.SearchGroup(department, groupname)
	if err != nil {
		zap.L().Error("dao.SearchGroup err", zap.Error(err))
		respond.Fail(c, respond.CodeSystemBusy)
		return
	}
	//如果找不到不会报错 所以再来一个判断
	if msg == nil && err == nil {
		respond.Fail(c, respond.CodeHasNotExit)
		return
	}
	respond.SuccessWith(c, msg)
}

// Add 加入用户信息
func Add(c *gin.Context) {
	msg := new(model.AllMsg)
	err := c.ShouldBindJSON(msg)
	if err != nil {
		respond.Fail(c, respond.CodeParamInvalid)
		return
	}

	department := c.GetString("department")
	err = logic.AddMember(department, msg)
	if err != nil {
		zap.L().Error("logic.AddMember err", zap.Error(err))
		respond.Fail(c, respond.CodeSystemBusy)
		return
	}
	respond.Success(c)
}

// DeleteData 删除信息
func DeleteData(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		respond.Fail(c, respond.CodeMsgNotReceive)
		return
	}

	department := c.GetString("department")
	err := logic.Delete(department, id)
	if err != nil {
		if errors.Is(err, model.ErrorAdminNotExit) {
			respond.Fail(c, respond.CodeHasNotExit)
			return
		}
		respond.Fail(c, respond.CodeSystemBusy)
		zap.L().Error("logic.Delete err", zap.Error(err))
		return
	}
	respond.Success(c)
}
