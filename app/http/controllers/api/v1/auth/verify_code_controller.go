package auth

import (
	v1 "forum/app/http/controllers/api"
	"forum/app/requests"
	"forum/pkg/captcha"
	"forum/pkg/logger"
	"forum/pkg/response"
	"forum/pkg/verifycode"
	"github.com/gin-gonic/gin"
)

// VerifyCodeController 基础控制器
type VerifyCodeController struct {
	v1.BaseAPIController
}

func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	// Create Captcha
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	// 记录错误日志
	logger.LogIf(err)
	// return user
	response.JSON(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}

func (vc *VerifyCodeController) SendUsingPhone(c *gin.Context) {

	// 1. 验证表单
	request := requests.VerifyCodePhoneRequest{}
	if ok := requests.Validate(c, &request, requests.VerifyCodePhone); !ok {
		return
	}

	// 2. 发送 SMS
	if ok := verifycode.NewVerifyCode().SendSMS(request.Phone); !ok {
		response.Abort500(c, "发送短信失败")
	} else {
		response.Success(c)
	}
}
