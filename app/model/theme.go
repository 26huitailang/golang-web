package model

import "github.com/jinzhu/gorm"

type Theme struct {
	gorm.Model
	Name   string `gorm:"not null;unique"`
	Suites []Suite
}
