package utils

import "fmt"

// 关注业务
const (
	PrefixFollowee = "followee" // 关注者
	PrefixFollower = "follower" // 粉丝
)

// followee:userId:bizId -> zset(objId, now)
func GetFolloweeKey(userId int64, bizId string) string {
	return fmt.Sprintf("%s:%d:%s", PrefixFollowee, userId, bizId)
}

// followee:userId:bizId -> zset(objId, now)
func GetFollowerKey(bizId string, userId int64) string {
	return fmt.Sprintf("%s:%d:%s", PrefixFollowee, userId, bizId)
}

// 点赞业务
const (
	PrefixEntityLike = "like:entity"
)

// like:entiey:bizId:objId -> set(userId)
func GetEntityLikeKey(bizId string, objId int64) string {
	return fmt.Sprintf("%s:%s:%d", PrefixEntityLike, bizId, objId)
}

// 收藏业务
const (
	PrefixCollect = "collect"
)

// collect:userId:bizId ->zet(objId,now)
func GetEntityCollectKey(userId int64, bizId string) string {
	return fmt.Sprintf("%s:%d:%s", PrefixCollect, userId, bizId)
}
