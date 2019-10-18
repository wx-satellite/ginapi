package main

import (
	"fmt"
	"gin/config"
	"gin/model"
	"gin/pkg/cache"
	"gin/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
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

	g.GET("/test1", func(context *gin.Context) {
		fmt.Println("/test1")
	})


	err = http.ListenAndServe(viper.GetString("addr"),g)

	fmt.Println(err)

}
