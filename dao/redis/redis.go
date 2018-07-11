package redis

import (
	"github.com/CX1ng/proxypool/common"
	"github.com/fzzy/radix/redis"
)

var redisHandler *redis.Client

func initRedisStorage(cfg *common.RedisConfig) {
	var err error
	redisHandler, err = redis.DialTimeout(cfg.Protocol, cfg.Dsn, cfg.Timeout.D())
	if err != nil {
		panic(err)
	}
}

func GetRedisHandler() *redis.Client {
	if redisHandler == nil {
		initRedisStorage(common.GetConfigHandler().Redis)
	}
	return redisHandler
}
