package initialize

import (
	"GinRESTful/restapi/global"
	"fmt"
	"github.com/go-redis/redis"
)

func InitRedis() {
	redisInfo := global.Settings.RedisInfo
	redisAddr := fmt.Sprintf("%s:%d",
		redisInfo.Host,
		redisInfo.Port,
	)
	// 生成Redis客户端
	global.Redis = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisInfo.Password,
		DB:       0,
	})
	// 连接Redis
	_, err := global.Redis.Ping().Result()
	if err != nil {
		panic(err)
	}
}
