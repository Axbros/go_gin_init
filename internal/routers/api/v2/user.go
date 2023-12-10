package v2

import (
	"gin_init/global"
	"gin_init/internal/service"
	"gin_init/pkg/ErrorCode"
	"gin_init/pkg/app"
	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() User {
	return User{}
}
func (u User) Login(c *gin.Context) {
	param := service.UserLoginRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(ErrorCode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	result, err := svc.UserLogin(&service.UserLoginRequest{Username: param.Username, Password: param.Password})
	if err != nil {
		global.Logger.Errorf("svc.UserLogin err: %v", err)
		response.ToErrorResponse(ErrorCode.ErrorCountTagFail)
		return
	}
	if result {
		token, _ := app.GenerateToken(param.Username, param.Password)
		response.ToResponse(200, "ok", token)
		return
	} else {
		response.ToResponse(301, ErrorCode.ErrorUsernameOrPassword.Msg(), "")
		return
	}

}
