package main

import (
	"fmt"

	"forum/bootstrap"
	btsConfig "forum/config"

	"github.com/gin-gonic/gin"
)

func init() {
	btsConfig.Initialize()
}

func main() {

	// new 一个 Gin Engine 实例
	router := gin.New()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	// 运行服务
	err := router.Run(":3000")
	if err != nil {
		// 错误处理
		fmt.Println(err.Error())
	}
}
