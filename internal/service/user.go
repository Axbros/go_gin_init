package service

type UserLoginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (svc *Service) UserLogin(param *UserLoginRequest) (bool, error) {
	return svc.dao.UserLogin(param.Username, param.Password)
}
