package routes

import (
	"forum/app/http/controllers/api/v1/auth"
	"github.com/gin-gonic/gin"
)

func RegisterAPIRouters(r *gin.Engine) {

	v1 := r.Group("v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
		}
	}
}
