package token
/* jwt 相关 */

import (
	"fmt"
	"gin/pkg/errno"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
)


type Context struct {
	Id uint64
	Username string
}

func Sign(ctx *gin.Context, c Context, secret string) (j string ,err error) {
	var (
		token *jwt.Token
	)
	// 加密盐
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": c.Id,
		"username": c.Username,
		// iat签发时间 nbf生效时间
		"nbf": time.Now().Unix(),
		"iat": time.Now().Unix(),
	})

	j, err = token.SignedString([]byte(secret))

	return
}


func ParseRequest(ctx *gin.Context) (*Context, error) {
	var (
		header string
		secret string
		j string
	)
	header = ctx.Request.Header.Get("Authorization")

	secret = viper.GetString("jwt_secret")

	if len(header) == 0 {
		return &Context{}, errno.AuthValidation
	}

	_, _ = fmt.Sscanf(header, "Bearer %s", &j)

	return Parse(j,secret)
}


func Parse(token string, secret string) (*Context, error) {
	var (
		content *Context
		err error
		t *jwt.Token
		claims jwt.MapClaims
		ok bool
	)

	content = &Context{}

	t, err = jwt.Parse(token, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errno.ErrJWTIsNotLegal
		}
		return []byte(secret), nil
	})

	if err != nil {
		return content, err
	}

	if claims, ok = t.Claims.(jwt.MapClaims); ok && t.Valid {
		// 需要拥float64断言，再强制转换
		content.Id = uint64(claims["id"].(float64))
		content.Username = claims["username"].(string)
		return content, nil
	}


	return content, errno.ErrJWTIsNotLegal


}

