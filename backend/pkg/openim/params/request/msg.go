package request

// MsgInfo msg info.
type MsgInfo struct {
	SendID           string           `json:"sendID"           binding:"required"`
	RecvID           string           `json:"recvID"           binding:"required"`
	GroupID          string           `json:"groupID"`
	SenderNickname   string           `json:"senderNickname"`
	SenderFaceURL    string           `json:"senderFaceURL"`
	SenderPlatformID int              `json:"senderPlatformID"`
	Content          *TextContent     `json:"content"          binding:"required"`
	ContentType      int              `json:"contentType"      binding:"required"`
	SessionType      int              `json:"sessionType"      binding:"required"`
	IsOnlineOnly     bool             `json:"isOnlineOnly"`
	NotOfflinePush   bool             `json:"notOfflinePush"`
	OfflinePushInfo  *OfflinePushInfo `json:"offlinePushInfo"`
}

// TextContent text content.
type TextContent struct {
	Text string `json:"content"`
}

// OfflinePushInfo offline push info.
type OfflinePushInfo struct {
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Ex            string `json:"ex"`
	IOSPushSound  string `json:"iOSPushSound"`
	IOSBadgeCount bool   `json:"iOSBadgeCount"`
}
