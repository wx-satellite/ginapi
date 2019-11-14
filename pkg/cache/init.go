package cache

import (
	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
	"time"
)

type ManagerCache struct {
	// 多个缓存，可扩展
	Redis *redis.Pool
}


var Cache *ManagerCache


func  Init() {
	// 初始化Cache
	Cache = &ManagerCache{}

	Cache.RedisInit()
}

func  Close() {
	Cache.RedisClose()
}

func (cache *ManagerCache) RedisInit() {

	cache.Redis = GetRedisExample()

}

func (cache *ManagerCache) RedisClose() {
	var (
		err error
	)
	if err = cache.Redis.Close(); err != nil {
		panic(err)
	}
}

func GetRedisExample() *redis.Pool {
	var (
		r *redis.Pool
	)
	r = &redis.Pool{
		MaxActive: 1000,
		MaxIdle: 30,
		IdleTimeout: 30 * time.Second,
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp",
				viper.GetString("redis.address"),
				redis.DialPassword(viper.GetString("redis.password")),
				redis.DialDatabase(viper.GetInt("redis.database")),
			)
		},
	}
	return r
}




