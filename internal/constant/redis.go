package constant

import "time"

const (
	UserCachePrefix  = "user:"
	UserListCacheKey = "user:list"
	UserCacheTTL     = 5 * time.Minute
	UserListCacheTTL = 2 * time.Minute
)
