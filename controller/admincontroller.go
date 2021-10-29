package controller

import (
	"bluebull/dao"
	"bluebull/logic"
	"bluebull/model"
	"bluebull/respond"
	"database/sql"
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
		respond.Fail(c, respond.CodeMsgReceive)
		return
	}

	department := c.GetString("department")
	Msg, err := dao.ShowAData(department, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respond.Fail(c, respond.CodeHasNotExit)
			return
		}
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
		respond.Fail(c, respond.CodeMsgReceive)
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
		if errors.Is(err, model.ErrorNotExit) {
			respond.Fail(c, respond.CodeHasNotExit)
			return
		}
		zap.L().Error("dao.Update err", zap.Error(err))
		respond.Fail(c, respond.CodeSystemBusy)
		return
	}
	respond.Success(c)
}
