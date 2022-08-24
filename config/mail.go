// Package config 站点配置信息
package config

import "forum/pkg/config"

func init() {
	config.Add("mail", func() map[string]interface{} {
		return map[string]interface{}{

			// 默认是腾讯企业邮箱的配置
			"smtp": map[string]interface{}{
				"host":     config.Env("MAIL_HOST", "smtp.exmail.qq.com"),
				"port":     config.Env("MAIL_PORT", 465),
				"username": config.Env("MAIL_USERNAME", ""),
				"password": config.Env("MAIL_PASSWORD", ""),
			},

			"from": map[string]interface{}{
				"address": config.Env("MAIL_FROM_ADDRESS", "gohub@example.com"),
				"name":    config.Env("MAIL_FROM_NAME", "Gohub"),
			},
		}
	})
}
