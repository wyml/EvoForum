package v1

import (
	v1 "forum/app/http/controllers/api"
	"forum/app/models/user"
	"forum/pkg/auth"
	"forum/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	v1.BaseAPIController
}

func (u *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}

func (u *UsersController) Index(c *gin.Context) {
	data := user.All()
	response.Data(c, data)
}
