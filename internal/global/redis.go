package global

import (
	redisTemp "github.com/galaxy-toolkit/server/database/redis"
	"github.com/redis/go-redis/v9"
)

// Redis 数据库链接
var Redis *redis.Client

// InitRedis 初始化 Redis 数据库
func InitRedis() {
	Redis = redisTemp.New(Config.Database.Redis)
}
