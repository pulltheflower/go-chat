package models

import (
	"errors"
	"fmt"
	"go-chat/initializer"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string `json:"name"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	ClientIp   string `json:"client_ip"`
	ClientPort string `json:"client_port"`
	LoginTime  uint64 `json:"login_time"`
	Heartbeat  uint64 `json:"heartbeat"`
	LogoutTime uint64 `json:"logout_time"`
	Offline    bool   `json:"offline"`
	DeviceInfo string `json:"device_info"`
}

type CreateUserAttributes struct {
	Name                 string `json:"name"`
	Password             string `json:"password"`
	PasswordConfirmation string
}

var (
	ErrPasswordMismatch = errors.New("The password is different from the confirmation password")
)

func (table *User) TableName() string {
	return "users"
}

func Users() []User {
	var users []User
	initializer.DB.Find(&users)
	for _, v := range users {
		fmt.Println(v)
	}
	return users
}

func CreateUser(user User) *gorm.DB {
	return initializer.DB.Create(&user)
}

func DeleteUser(user User) *gorm.DB {
	return initializer.DB.Delete(&user)
}

func UpdateUser(user User) *gorm.DB {
	return initializer.DB.Model(&user).Updates(User{Name: user.Name, Password: user.Password})
}

func (c CreateUserAttributes) Validation() error {
	switch {
	case c.Password != c.PasswordConfirmation:
		return ErrPasswordMismatch
	default:
		return nil
	}
}
