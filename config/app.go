package config

import "forum/pkg/config"

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			"name":     config.Env("APP_NAME", "MLTREE FORUM"),                                                    // 应用名称
			"env":      config.Env("APP_ENV", "production"),                                                       // 当前环境，local, stage, production, test
			"debug":    config.Env("APP_DEBUG", false),                                                            // 调试模式
			"port":     config.Env("APP_PORT", "8080"),                                                            // 服务端口
			"key":      config.Env("APP_KEY", "LUpSgnM3pzGnhfKlrd1NTKkY79v3EQxp6QBqAaq21cBNUVD1qHBGvi8DEYO9b2RH"), // 加密密钥
			"url":      config.Env("APP_URL", "http://localhost:8080"),                                            // 服务地址
			"timezone": config.Env("TIMEZON", "Asia/Shanghai"),                                                    // 时区
		}
	})
}
