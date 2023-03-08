package utils

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/redis/go-redis/v9"
)

type DbConfig struct {
	Redis1  redis.Client
	Redis5  redis.Client
	MysqlDb gdb.DB
}

var G *DbConfig

func Global() {
	G = new(DbConfig)
	G.Redis1 = RedisInit(1)
	G.Redis5 = RedisInit(5)
	G.MysqlDb = MysqlInit(true)
}
