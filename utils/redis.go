package utils

import (
	conf "EasyGinLite/config"
	"fmt"
	"github.com/redis/go-redis/v9"
)

// RedisInit 初始化一个Redis的实例，使用参考：https://redis.uptrace.dev/zh/guide/go-redis.html
func RedisInit(db int) (rdb redis.Client) {
	rdb = *redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf(`%s:%d`, conf.ProjectConfig.Redis.Addr, conf.ProjectConfig.Redis.Port),
		Password: conf.ProjectConfig.Redis.Password,
		DB:       db,
	})
	return
}
