package config

import "forum/pkg/config"

func init() {
	config.Add("verifycode", func() map[string]interface{} {
		return map[string]interface{}{
			"code_length": config.Env("VERIFY_CODE_LENGTH", 6),  // 验证码长度
			"expire_time": config.Env("VERIFY_CODE_EXPIRE", 15), // 过期时间，单位：分钟

			// 调试模式下参数
			"debug_expire_time":  10080,          // 调试模式过期时间
			"debug_code":         123456,         // 调试模式验证码
			"debug_phone_prefix": "000",          // 调试模式手机号前缀
			"debug_email_suffix": "@testing.com", // 调试模式邮箱后缀
		}
	})
}
