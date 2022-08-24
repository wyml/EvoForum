package middlewares

import (
	"forum/pkg/jwt"
	"forum/pkg/response"
	"github.com/gin-gonic/gin"
)

// GuestJWT 强制使用游客身份访问，如用户注册、登录接口
func GuestJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) > 0 {
			_, err := jwt.NewJWT().ParserToken(c)
			if err == nil {
				response.Unauthorized(c, "请使用游客身份访问")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
