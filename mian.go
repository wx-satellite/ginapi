package main

import (
	"errors"
	"fmt"
	"gin/config"
	"gin/model"
	"gin/pkg/cache"
	"gin/router"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func main() {
	var (
		err error
		middlewares []gin.HandlerFunc
		g *gin.Engine
	)

	// 加载配置文件
	if err = config.Init(""); err != nil {
		panic(err)
	}



	// 初始化数据库
	model.Init()
	defer model.Close()


	// 初始化缓存
	cache.Init()
	defer cache.Close()


	gin.SetMode(viper.GetString("runmode"))

	g = gin.New()



	middlewares = []gin.HandlerFunc{}

	router.Load(
		g,
		middlewares...,
	)

	//心跳检测
	go func() {
		if err = pingService(); err != nil {
			log.Fatalf(err,"心跳检测失败，路由不可用！错误：%v", err)
		}
		log.Info("心跳检测成功，路由可用。")
	}()

	g.GET("/test1", func(context *gin.Context) {
		context.String(http.StatusOK, "it is a test")
	})


	err = http.ListenAndServe(viper.GetString("addr"),g)

	fmt.Println(err)
}

// api进程运行成功不代表api服务可以正常对外提供服务，因此做了自检
func pingService() error {
	var (
		addr string
		response *http.Response
		err error
		retry int
		password string
		url string
	)
	addr = viper.GetString("addr")
	retry = viper.GetInt("health_retry_count")
	password = viper.GetString("health_password")
	url = fmt.Sprintf("http://127.0.0.1%s/health/ping?password=%s", addr, password)
	for i := 1; i <= retry ; i++ {
		response, err = http.Get(url)
		if response != nil && err == nil && response.StatusCode == http.StatusOK {
			return nil
		}
		log.Infof("尝试在1秒之后重试！")
		time.Sleep(1 * time.Second)
	}
	return errors.New("心跳检测失败！")
}