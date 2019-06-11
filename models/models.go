package models

import "github.com/jinzhu/gorm"

type Theme struct {
	gorm.Model
	Name   string `gorm:"not null;unique"`
	Suites []Suite
}

type Suite struct {
	gorm.Model
	Name    string `gorm:"size:300;not null;unique"`
	ThemeID uint
	Images  []Image
	IsRead  bool `gorm:"DEFAULT:false"`
}

type Image struct {
	gorm.Model
	Path    string
	SuiteID uint
}
