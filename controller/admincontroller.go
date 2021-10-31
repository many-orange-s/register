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
	if msg == nil && err == nil {
		respond.Fail(c, respond.CodeHasNotExit)
		return
	}
	respond.SuccessWith(c, msg)
}

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
