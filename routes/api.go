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
			// 判断邮箱是否已注册
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
			// 使用手机号注册
			authGroup.POST("/signup/using-phone", suc.SignupUsingPhone)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			// 图片验证码，需要加限流
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
			// 短信验证码
			authGroup.POST("/verify-codes/phone", vcc.SendUsingPhone)
			// 邮箱验证码
			authGroup.POST("/verify-codes/email", vcc.SendUsingEmail)
		}
	}
}
