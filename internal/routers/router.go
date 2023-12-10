package routers

import (
	"gin_init/global"
	"gin_init/internal/middleware"
	"gin_init/internal/routers/api"
	v1 "gin_init/internal/routers/api/v1"
	v2 "gin_init/internal/routers/api/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.POST("/auth", api.GetAuth)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	tag := v1.NewTag()
	userHandler := v2.NewUser()

	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.JWT())
	{
		apiV1.POST("/tags", tag.Create)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PUT("/tags/:id", tag.Update)
		apiV1.PATCH("/tags/:id/state", tag.Update)
		apiV1.GET("/tags", tag.List)
	}
	apiV2 := r.Group("/api/v2") //v2 doesn't need authority
	{
		apiV2.POST("/login", userHandler.Login)
	}

	return r
}
