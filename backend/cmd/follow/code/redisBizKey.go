package code

import "fmt"

func UserFollowKey(userId string) string {
	return fmt.Sprintf("biz:user:follow:%s", userId)
}

func UserFansKey(userId string) string {
	return fmt.Sprintf("biz:user:fans:%s", userId)
}
