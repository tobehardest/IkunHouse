package code

import "fmt"

func UserFollowKey(userId int64) string {
	return fmt.Sprintf("biz:user:follow:%d", userId)
}

func UserFansKey(userId int64) string {
	return fmt.Sprintf("biz:user:fans:%d", userId)
}
