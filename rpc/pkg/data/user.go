package data

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Id           int64  `gorm:"column:id" json:"id"`
	Login        string `gorm:"column:login" json:"login"`
	Email        string `gorm:"column:email" json:"email"`
	Password     string `gorm:"column:password" json:"password"`
	Role         string `gorm:"column:role" json:"role" gorm:"default:'user'"`
	RefreshToken string `gorm:"column:refresh_token" json:"refresh_token"`
}

type UserData interface {
	Get(login, password string) (*User, error)
	Create(user *User) (int64, error)
	GetByLogin(login string) (*User, error)
	Update(user *User, refreshToken string) (*User, error)
}

type userData struct {
	db *gorm.DB
}

func NewUserData(db *gorm.DB) *userData {
	return &userData{db}
}

func (d *userData) Get(login, password string) (*User, error) {
	user := User{}
	if err := d.db.Where("login = ? AND password = ?", login, password).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (d *userData) GetByLogin(login string) (*User, error) {
	user := User{}
	if err := d.db.Where("login = ?", login).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (d *userData) Create(user *User) (int64, error) {
	if err := d.db.Create(user).Error; err != nil {
		return -1, err
	}
	return user.Id, nil
}

func (d *userData) Update(user *User, refreshToken string) (*User, error) {
	if err := d.db.Model(&user).Where("id = ?", user.Id).Update("refresh_token", refreshToken).Error; err != nil {
		return nil, err
	}
	return user, nil
}
