package model

import "github.com/jinzhu/gorm"

type Image struct {
	gorm.Model
	Path    string
	SuiteID uint
}
