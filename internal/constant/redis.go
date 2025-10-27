package constant

import "time"

const (
	UserCachePrefix  = "user:"
	UserListCacheKey = "user:list"
	UserCacheTTL     = 5 * time.Minute
	UserListCacheTTL = 2 * time.Minute
)

const (
	ErrRedisGetUserMsg  = "faild to get user from cache"
	ErrRedisListuserMsg = "faild to get list user from cache"
)
