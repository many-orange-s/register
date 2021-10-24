package respond
type ReCode int64

const (
	CodeSuccess ReCode = 1000 + iota
	CodeParamInvalid
	CodeParamFalse
	CodeSystemBusy
)

var codeMsgMap = map[ReCode]string{
	CodeSuccess:      "Success",
	CodeParamInvalid: "Invalid params",
	CodeParamFalse:   "The account number or password is incorrect",
	CodeSystemBusy:   "The system is busy",
}

type RespondDate struct {
	Code ReCode      `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}
