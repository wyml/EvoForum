package paginator

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Paging struct {
	CurrentPage int
	PerPage     int
	TotalPage   int
	TotalCount  int64
	NextPageURL string
	PrevPageURL string
}

type Paginator struct {
	BaseURL    string
	PerPage    int
	Page       int
	Offset     int
	TotalCount int64
	TotalPage  int
	Sort       string
	Order      string

	query *gorm.DB
	ctx   *gin.Context
}
