package router

import (
	"gin/router/middleware"
	"github.com/gin-gonic/gin"
)

func Load(engine *gin.Engine, mw ...gin.HandlerFunc) {
	var (
		uGroup *gin.RouterGroup
		sGroup *gin.RouterGroup
		v1Edition *gin.RouterGroup
	)
	// 处理某些请求时可能因为bug或者异常导致程序panic，为了不影响下一次请求的调用，需要通过gin.Recover()来恢复服务器
	engine.Use(gin.Recovery())
	// 强制浏览器不使用缓存
	engine.Use(middleware.NoCache)
	// 跨域设置
	engine.Use(middleware.Options)
	// 自定义的中间件
	engine.Use(mw...)


	// v1版本
	v1Edition = engine.Group("/v1")
	{
		// 会话相关
		sGroup = v1Edition.Group("session")
		{
			sGroup.GET("/s", func(context *gin.Context) {
				context.JSON(200, "it is a test~")
			})
		}

		// 用户相关
		uGroup = v1Edition.Group("user")
		{
			uGroup.GET("/u", func(context *gin.Context) {
				context.JSON(200,"it is a test")
			})
		}

	}




}
