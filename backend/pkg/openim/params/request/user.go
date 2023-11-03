package request

// RegisterUserParams register user params.
type RegisterUserParams struct {
	Secret string `json:"secret" binding:"required"`
	Users  []User `json:"users"  binding:"required"`
}

// User user.
type User struct {
	UserID   string `json:"userID"   binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	FaceURL  string `json:"faceURL"  binding:"required"`
}
