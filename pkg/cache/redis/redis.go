package redis

import (
	"gin/pkg/cache"
	"github.com/garyburd/redigo/redis"
)


func Set(key string, value interface{}) (err error) {
	_, err = cache.Cache.Redis.Do("SET",key, value)
	return
}

func Setex(key string, value string, second uint64) (err error) {
	_, err = cache.Cache.Redis.Do("SET", key, value, "EX", second)
	return
}


func Setnx(key string, value string) (success bool, err error) {
	n, err := cache.Cache.Redis.Do("SETNX", key, value)
	if err != nil {
		return false, err
	}
	if int64(1) == n {
		return true, nil
	} else {
		return false, nil
	}

}


func Get(key string) (string, error) {
	return redis.String(cache.Cache.Redis.Do("GET", key))
}

func Del(key string) (err error) {
	_, err = cache.Cache.Redis.Do("DEL", key)
	return
}




