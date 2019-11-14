package redis

import (
	"gin/pkg/cache"
	"github.com/garyburd/redigo/redis"
)


func Set(key string, value interface{}) (err error) {
	con := cache.Cache.Redis.Get()
	defer con.Close()
	_, err = con.Do("SET",key, value)
	return
}

func Get(key string) (string, error) {
	con := cache.Cache.Redis.Get()
	defer con.Close()
	return redis.String(con.Do("GET", key))
}

func Del(key string) (err error) {
	con := cache.Cache.Redis.Get()
	defer con.Close()
	_, err = con.Do("DEL", key)
	return
}

//func Setex(key string, value string, second uint64) (err error) {
//	_, err = cache.Cache.Redis.Do("SET", key, value, "EX", second)
//	return
//}
//
//
//func Setnx(key string, value string) (success bool, err error) {
//	n, err := cache.Cache.Redis.Do("SETNX", key, value)
//	if err != nil {
//		return false, err
//	}
//	if int64(1) == n {
//		return true, nil
//	} else {
//		return false, nil
//	}
//
//}







