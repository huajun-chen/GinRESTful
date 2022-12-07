package utils

import (
	"GinRESTful/restapi/global"
	"time"
)

// RedisSetStr Redis Set 字符串
// 参数：
//		key：Redis的key
//		value：Redis的value
//		expiration：Redis的key的过期时间
// 返回值：
//		error：错误信息
func RedisSetStr(key string, value interface{}, expiration time.Duration) error {
	err := global.Redis.Set(key, value, expiration).Err()
	return err
}

// RedisGetStr Redis Get 字符串
// 参数：
//		key：Redis的key
// 返回值：
//		string：获取的key对应的value值
//		error：错误信息
func RedisGetStr(key string) (string, error) {
	value, err := global.Redis.Get(key).Result()
	return value, err
}
