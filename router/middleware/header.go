package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func NoCache(ctx *gin.Context) {
	ctx.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	ctx.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	ctx.Header("Last-Modified",time.Now().UTC().Format(http.TimeFormat))
	ctx.Next()
}

func Options(ctx *gin.Context) {
	if ctx.Request.Method != "OPTIONS" {
		ctx.Next()
	} else {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods","GET,POST,PUT,PATCH,DELETE,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers","authorization, origin, content-type, accept")
		ctx.Header("Allow","HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		ctx.Header("Content-Type", "application/json")
		ctx.AbortWithStatus(200)
	}
}
