package cache

import (
	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
)

type ManagerCache struct {
	// 多个缓存，可扩展
	Redis redis.Conn
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

func GetRedisExample() redis.Conn {
	var (
		err error
		r redis.Conn
	)
	r, err = redis.Dial("tcp",
		viper.GetString("redis.address"),
		redis.DialPassword(viper.GetString("redis.password")),
		redis.DialDatabase(viper.GetInt("redis.database")),
	)
	if err != nil {
		panic(err)
	}
	return r
}




