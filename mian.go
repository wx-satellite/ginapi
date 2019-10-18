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

/*
 一些规范：变量的声明最好在方法一开始就使用var声明而不是使用":="这种方式。这样做有一个好处就是配合"goto"使用很方便
var (
	name string
	age int
)

 */

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

	// Routes.
	router.Load(
		// Cores.
		g,

		middlewares...,
	)

	g.GET("/test1", func(context *gin.Context) {
		fmt.Println("/test1")
	})


	err = http.ListenAndServe(viper.GetString("addr"),g)

	fmt.Println(err)

}
