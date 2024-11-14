package model

type LResp struct {
	Code int
	Msg  string
	Data any
}

type LError struct {
	LResp
}

func NewErr(code int, msg string, data any) *LError {
	return &LError{
		LResp: LResp{Msg: msg, Code: code, Data: data},
	}
}

func LErr2Json(err *LError) map[string]any {
	return respFormat(err.Code, err.Msg, err.Data)
}

func Json(code int, msg string, data any) map[string]any {
	return respFormat(code, msg, data)
}

func JsonSuc(msg string, data any) map[string]any {
	return respFormat(200, msg, data)
}

func respFormat(code int, msg string, data any) map[string]any {
	return map[string]any{"code": code, "msg": msg, "data": data}
}
