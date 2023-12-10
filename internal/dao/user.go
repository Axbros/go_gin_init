package dao

import "gin_init/internal/model"

func (d *Dao) UserLogin(name, password string) (bool, error) {
	user := model.User{Username: name, Password: password, IsAdmin: 0}
	return user.Login(d.engine)
}
