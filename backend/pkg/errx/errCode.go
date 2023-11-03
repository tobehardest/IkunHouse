package errx

// 成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

// 全局错误码
const SERVER_COMMON_ERROR uint32 = 100001
const REUQEST_PARAM_ERROR uint32 = 100002
const TOKEN_EXPIRE_ERROR uint32 = 100003
const TOKEN_GENERATE_ERROR uint32 = 100004
const DB_ERROR uint32 = 100005
const DB_UPDATE_AFFECTED_ZERO_ERROR uint32 = 100006
<<<<<<< HEAD

//用户模块

// 关注模块
const (
	FollowUserIdEmpty   = 400001
	FollowedUserIdEmpty = 400002
	CannotFollowSelf    = 400003
	UserIdEmpty         = 400004
)

// 视频模块
=======
>>>>>>> yikun-dev
