package models

type ResposeCode int64

const (
	CodeSuccess ResposeCode = 1000

	CodeInvalidParams ResposeCode = 1001
	CodeTokenError    ResposeCode = 1002
	CodeServerBusy    ResposeCode = 1003
)

var msgFlags = map[ResposeCode]string{
	CodeSuccess:       "success",
	CodeInvalidParams: "参数错误",
	CodeTokenError:    "TOKEN过期，请重新登陆",
	CodeServerBusy:   "服务器繁忙",
}

func (c ResposeCode) Msg() string {
	msg, ok := msgFlags[c]
	if ok {
		return msg
	}
	return msgFlags[CodeServerBusy]
}
