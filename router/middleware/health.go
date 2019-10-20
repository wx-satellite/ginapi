package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func CanCheckHealth(ctx *gin.Context){
	var (
		password string
	)

	password = ctx.Query("password")


	if password != viper.GetString("health_password") {
		ctx.String(http.StatusUnauthorized, "你没有权限！")
		ctx.Abort()
		return
	}
	ctx.Next()
}