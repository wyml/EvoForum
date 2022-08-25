package v1

import (
	v1 "forum/app/http/controllers/api"
	"forum/app/models/user"
	"forum/app/requests"
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
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}
	
	data, pager := user.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}
