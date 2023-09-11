package globalKey

var (
	DefaultUserHeaderImgID int64 = 16 // 用户默认头像ID
	UserStateSites         int64 = 1  // 用户网站注册
	UserStateQQ            int64 = 2  // 用户qq注册
	UserStateWX            int64 = 3  // 用户微信注册
	UserRoleBan            int64 = -1 // 被禁用用户
	UserRoleComm           int64 = 1  // 普通用户
	UserRoleAdmin          int64 = 2  // 管理员用户
	UserNotBoss            int64 = 0  // 不是商家
	UserIsBoss             int64 = 1  // 商家
)
