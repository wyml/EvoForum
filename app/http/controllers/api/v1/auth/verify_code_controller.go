package auth

import (
	v1 "forum/app/http/controllers/api"
	"forum/pkg/captcha"
	"forum/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
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
	c.JSON(http.StatusOK, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}
