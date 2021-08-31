package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ApiLoginReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Session struct {
	gorm.Model
	Token     string `json:"token" gorm:"index,unique"`
	Value     string `json:"value"`
	ExpiredAt time.Time
}

type SessionValue struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}
