package auth

import (
	"fmt"
	v1 "forum/app/http/controllers/api"
	"forum/app/models/user"
	"forum/app/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist 检测手机号是否已注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {

	// 初始化请求对象
	request := requests.SignupPhoneExistRequest{}

	// 解析 JSON 请求
	if err := c.ShouldBindJSON(&request); err != nil {
		// 解析失败，返回 422 状态码以及错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		// Print Error Message
		fmt.Println(err.Error())
		// Exit request
		return
	}

	// Validate
	errs := requests.ValidateSignupPhoneExist(&request, c)
	if len(errs) > 0 {
		// 验证失败，返回 422 状态码以及错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	c.JSONP(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

// IsEmailExist 检测邮箱是否已注册
func (sc *SignupController) IsEmailExist(c *gin.Context) {

	// 请求对象
	type PhoneExistRequest struct {
		Email string `json:"email"`
	}
	request := requests.SignupEmailExistRequest{}

	// 解析 JSON 请求
	if err := c.ShouldBindJSON(&request); err != nil {
		// 解析失败，返回 422 状态码以及错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		// Print Error Message
		fmt.Println(err.Error())
		// Exit request
		return
	}

	// Validate
	errs := requests.ValidateSignupEmailExist(&request, c)
	if len(errs) > 0 {
		// 验证失败，返回 422 状态码以及错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	c.JSONP(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Email),
	})
}
