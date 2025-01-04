package output

import (
	"fmt"
)

type LResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type LError struct {
	LResp
}

func Err(code int, msg string, data any) *LError {
	return &LError{
		LResp: LResp{Code: code, Msg: msg, Data: data},
	}
}

func Suc(msg string, data any) *LResp {
	return &LResp{Code: 0, Msg: msg, Data: data}
}

func Err2Str(err *LError) string {
	return fmt.Sprintf("%+v", err)
}
