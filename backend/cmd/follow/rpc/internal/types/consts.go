package types

const (
	FollowStatusFollow   = iota + 1 // 关注
	FollowStatusUnfollow            // 取消关注
)

const (
	DefaultPageSize     = 20
	CacheMaxFollowCount = 1000 // 缓存最大关注数
	CacheMaxFansCount   = 1000 // 缓存最大粉丝数
)

const (
	UserFollowExpireTime = 3600 * 24 * 2
)
