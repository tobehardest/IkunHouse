package errorx

type BaseResponse struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}
