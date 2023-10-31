package request

// UserTokenParams register admin params.
type UserTokenParams struct {
	Secret     string `json:"secret"     binding:"required"`
	PlatformID uint   `json:"platformID" binding:"required"`
	UserID     string `json:"userID"     binding:"required"`
}
