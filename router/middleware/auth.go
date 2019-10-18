package middleware

import (
	"gin/handler"
	"gin/pkg/errno"
	"gin/pkg/token"
	"github.com/gin-gonic/gin"
)

func CheckAuth(ctx *gin.Context) {
	if content, err := token.ParseRequest(ctx); err != nil {
		handler.SendResponse(ctx, errno.ErrJWTIsNotLegal, nil)
		ctx.AbortWithStatus(401)
		return
	} else {
		ctx.Set("user", content)
		ctx.Next()
	}
}