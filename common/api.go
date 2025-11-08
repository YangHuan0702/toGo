package common

type ToGoResponse struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	Obj     interface{} `json:"obj"`
	Success bool        `json:"success"`
}

func Success(data interface{}) ToGoResponse {
	return ToGoResponse{Code: 200, Data: data, Success: true}
}

func Finish() ToGoResponse {
	return ToGoResponse{Code: 200, Success: true}
}

func Error(msg string) ToGoResponse {
	return ToGoResponse{Code: 500, Msg: msg, Success: false}
}

func Exception(msg string, code int) ToGoResponse {
	return ToGoResponse{Code: code, Msg: msg, Success: false}
}
