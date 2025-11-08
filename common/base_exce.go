package common

type ExceptionCode = int
type ApiException = *ToBeException

type ToBeException struct {
	Code ExceptionCode `json:"code"`
	Msg  string        `json:"msg"`
}

const NilParams ExceptionCode = 500
const NotFindMenu ExceptionCode = 1403
const NotFindParentMenu ExceptionCode = 1413

var ExceptionRespMap = map[ExceptionCode]ApiException{
	NilParams:         {Code: NilParams, Msg: "空值参数"},
	NotFindMenu:       {Code: NotFindMenu, Msg: "找不到目标菜单"},
	NotFindParentMenu: {Code: NotFindParentMenu, Msg: "找不到目标父级菜单"},
}
