package verifycode

import (
	"forum/pkg/app"
	"forum/pkg/config"
	"forum/pkg/helpers"
	"forum/pkg/logger"
	"forum/pkg/redis"
	"forum/pkg/sms"
	"strings"
	"sync"
)

type VerifyCode struct {
	Store Store
}

var once sync.Once
var internalVerifyCode *VerifyCode

func NewVerifyCode() *VerifyCode {
	once.Do(func() {
		internalVerifyCode = &VerifyCode{
			Store: &RedisStore{
				RedisClient: redis.Redis,
				KeyPrefix:   config.GetString("app.name") + ":verifycode",
			},
		}
	})
	return internalVerifyCode
}

// SendSMS 发送短信验证码，调用示例
// verifycode.NewVerifyCode().SendSMS(request.Phone)
func (vc *VerifyCode) SendSMS(phone string) bool {

	// Create code
	code := vc.generateVerifyCode(phone)

	// 方便本地
	if !app.IsProduction() && strings.HasPrefix(phone,
		config.GetString("verifycode.debug_phone_prefix")) {
		return true
	}
	return sms.NewSMS().Send(phone, sms.Message{
		Template: config.GetString("sms.aliyun.template_code"),
		Data:     map[string]string{"code": code},
	})
}

// CheckAnswer 检查用户提交的验证码是否正确， key 可以是手机号或者 Email
func (vc *VerifyCode) CheckAnswer(key string, answer string) bool {

	logger.DebugJSON("验证码", "检查验证码", map[string]string{key: answer})

	if !app.IsProduction() &&
		(strings.HasSuffix(key, config.GetString("verifycode.debug_email_suffix")) ||
			strings.HasPrefix(key, config.GetString("verifycode.debug_phone_prefix"))) {
		return true
	}

	return vc.Store.Verify(key, answer, false)
}

// 生产验证码，并放置于 Redis 中
func (vc *VerifyCode) generateVerifyCode(key string) string {

	// 生产随机码
	code := helpers.RandomNumber(config.GetInt("verifycode.code_length"))

	// 开发环境使用固定验证码
	if app.IsLocal() {
		code = config.GetString("verifycode.debug_code")
	}

	logger.DebugJSON("验证码", "生产验证码", map[string]string{key: code})

	// 将验证码以及 KEY 存放到 Redis 中并设置过期时间
	vc.Store.Set(key, code)
	return code
}
