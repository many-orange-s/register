package respond

/*里面只放返回参数*/

type ReCode int64

const (
	CodeSuccess ReCode = 1000 + iota
	CodeParamInvalid
	CodeParamFalse
	CodeSystemBusy
	CodeAdminExit
	CodeEmptyAuth
	CodeAuthFormatInvalid
	CodeTokenInvalid
	CodeMsgNotReceive
	CodeHasNotExit
	CodeRepeatModify
)

var codeMsgMap = map[ReCode]string{
	CodeParamInvalid:      "Invalid params",
	CodeParamFalse:        "The account number or password is incorrect",
	CodeSystemBusy:        "The system is busy",
	CodeAdminExit:         "The Admin has exited",
	CodeEmptyAuth:         "The Header's authorization is empty",
	CodeAuthFormatInvalid: "The Header's authorization format err",
	CodeTokenInvalid:      "Token is invalid",
	CodeSuccess:           "Success",
	CodeMsgNotReceive:     "Can receive this massage",
	CodeHasNotExit:        "Not exit",
	CodeRepeatModify:      "Repeat modify",
}

type RespondDate struct {
	Code ReCode      `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}
