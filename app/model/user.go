package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:100;not null;unique"`
	Password string `gorm:"not null;"`
	Nickname string `gorm:"size:100;not null;unique"`
}
