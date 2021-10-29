package controller

import (
	"bluebull/JWT/Token"
	"bluebull/dao"
	"bluebull/logic"
	"bluebull/model"
	"bluebull/respond"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AdminSign(c *gin.Context) {
	p := new(model.Sign)
	if err := c.ShouldBindJSON(p); err != nil {
		respond.Fail(c, respond.CodeParamInvalid)
		return
	}

	deparment, err := dao.SelectAdmin(p)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respond.Fail(c, respond.CodeParamFalse)
		} else {
			zap.L().Error("dao.SelectAdmin", zap.Error(err))
			respond.Fail(c, respond.CodeSystemBusy)
		}
		return
	}

	tokenstring, err := Token.GetToken(deparment)
	if err != nil {
		zap.L().Error("Token.GetToken err", zap.Error(err))
		respond.FailWithMsg(c, model.ErrorToken, 400)
		return
	}
	respond.SuccessWith(c, tokenstring)
}

func AddAdmin(c *gin.Context) {
	p := new(model.Register)
	if err := c.ShouldBind(p); err != nil {
		respond.Fail(c, respond.CodeParamInvalid)
		return
	}

	err := logic.Logon(p)
	if err != nil {
		if errors.Is(err, model.ErrorAdminExit) {
			respond.Fail(c, respond.CodeAdminExit)
		} else {
			zap.L().Error("logic.Logon err", zap.Error(err))
			respond.Fail(c, respond.CodeSystemBusy)
		}
		return
	}

	err = dao.CreateForm(p.Department)
	if err != nil {
		zap.L().Error("dao.CreateForm err", zap.Error(err))
		respond.Fail(c, respond.CodeSystemBusy)
		return
	}

	err = dao.SignUp(p)
	if err != nil {
		zap.L().Error("dao.SignUp err", zap.Error(err))
		respond.Fail(c, respond.CodeSystemBusy)
		return
	}
	respond.Success(c)
}
