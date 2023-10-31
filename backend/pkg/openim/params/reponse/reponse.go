package response

// BaseResponse base response.
type BaseResponse struct {
	ErrCode uint        `json:"errCode"`
	ErrMsg  string      `json:"errMsg"`
	ErrDlt  string      `json:"errDlt"`
	Data    interface{} `json:"data"`
}

// UserTokenResponse user token response.
type UserTokenResponse struct {
	BaseResponse
	Data TokenData `json:"data"`
}

// TokenData token data.
type TokenData struct {
	Token             string `json:"token"`
	ExpireTimeSeconds uint   `json:"expireTimeSeconds"`
}

// BotQueryResponse bot query response.
type BotQueryResponse struct {
	Code uint         `json:"code"`
	Msg  string       `json:"msg"`
	Data BotQueryData `json:"data"`
}

// BotQueryData bot query data.
type BotQueryData struct {
	Context  string `json:"context"`
	Question string `json:"question"`
	Text     string `json:"text"`
}
