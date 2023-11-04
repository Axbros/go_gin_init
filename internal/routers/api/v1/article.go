package v1

import (
	"gin_init/pkg/ErrorCode"
	"gin_init/pkg/app"
	"github.com/gin-gonic/gin"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(ErrorCode.ServerError)
	return
}
func (a Article) List(c *gin.Context)   {}
func (a Article) Create(c *gin.Context) {}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}
