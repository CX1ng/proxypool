package redis

import (
	"time"

	"github.com/fzzy/radix/redis"

	. "github.com/CX1ng/proxypool/common"
)

var testRedisConnector = RedisConnector{}

func init() {
	InitConfig("../../config/config.test.toml")
	testRedisConn, err := redis.DialTimeout("tcp", "127.0.0.1:6379", 3*time.Second)
	if err != nil {
		panic(err)
	}
	if testRedisConn == nil {
		panic("init redis conn failed")
	}
	testRedisConnector.conn = testRedisConn
}
