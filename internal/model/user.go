package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	*Model
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  int    `json:"is_admin"`
}

func (u User) TableName() string {
	return "axbros_user"
}
func (u User) Login(db *gorm.DB) (bool, error) {
	var user User
	if err := db.Where("username = ? AND password=? AND is_admin=?", u.Username, u.Password, u.IsAdmin).Find(&user).Error; err != nil {
		return false, nil
	}
	return true, nil

}
