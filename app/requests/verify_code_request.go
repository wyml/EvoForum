package requests

import (
	"forum/pkg/captcha"
	"forum/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type VerifyCodePhoneRequest struct {
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`

	Phone string `json:"phone,omitempty" valid:"phone"`
}

func VerifyCodePhone(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone":          []string{"required", "digits:11"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:" + config.GetString("verifycode.code_length")},
	}

	msg := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项",
			"digits:手机号必须为 11 位数字",
		},
		"captcha_id": []string{
			"required:图片验证码的 ID 为必填项",
		},
		"captcha_answer": []string{
			"required:图片验证码答案为必填项",
			"digits:图片验证码长度必须为 " + config.GetString("verifycode.code_length") + " 位数字",
		},
	}

	errs := validate(data, rules, msg)

	// 图片验证码
	_data := data.(*VerifyCodePhoneRequest)
	if ok := captcha.NewCaptcha().VerifyCaptcha(_data.CaptchaID, _data.CaptchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	}

	return errs
}
