package v1

import (
	v1 "forum/app/http/controllers/api"
	"forum/app/models/category"
	"forum/app/requests"
	"forum/pkg/response"

	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	v1.BaseAPIController
}

func (ctrl *CategoriesController) Index(c *gin.Context) {
	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := category.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

// Store 创建分类
func (ctrl *CategoriesController) Store(c *gin.Context) {

	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}

	categoryModel := category.Category{
		Name:        request.Name,
		Description: request.Description,
	}
	categoryModel.Create()
	if categoryModel.ID > 0 {
		response.Created(c, categoryModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *CategoriesController) Update(c *gin.Context) {

	// 验证 URL 参数 id 是否正确
	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID == 0 {
		response.Abort404(c)
		return
	}

	// 表单验证
	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}

	// 保存数据
	categoryModel.Name = request.Name
	categoryModel.Description = request.Description
	rowsAffected := categoryModel.Save()

	if rowsAffected > 0 {
		response.Data(c, categoryModel)
	} else {
		response.Abort500(c)
	}
}

func (ctrl *CategoriesController) Delete(c *gin.Context) {
	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID == 0 {
		response.Abort500(c)
		return
	}

	rowsAffected := categoryModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}