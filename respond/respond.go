package respond

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r ReCode) getMsg() string {
	msg,ok := codeMsgMap[r]
	if !ok{
		return codeMsgMap[CodeSystemBusy]
	}
	return msg
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK,&RespondDate{
		Code :CodeSuccess,
		Msg : CodeSuccess.getMsg(),
		Data: data,
	})
}

func Fail(c *gin.Context,code ReCode) {
	c.JSON(http.StatusOK,&RespondDate{
		Code: code,
		Msg: code.getMsg(),
		Data: nil,
	})
}

func FailWithMsg(c *gin.Context,msg interface{},code ReCode){
	c.JSON(http.StatusOK,&RespondDate{
		Code:code,
		Msg: msg,
		Data: nil,
	})
}


