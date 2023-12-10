package api

import (
	"gin_init/global"
	"gin_init/internal/service"
	"gin_init/pkg/ErrorCode"
	"gin_init/pkg/app"
	"gin_init/pkg/convert"
	"gin_init/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")

	if err != nil {
		response.ToErrorResponse(ErrorCode.InvalidParams.WithDetails(err.Error()))
		return
	}

	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(ErrorCode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Error(c, "svc.UploadFile err: %v", err)
		response.ToErrorResponse(ErrorCode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(200, "success", gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
