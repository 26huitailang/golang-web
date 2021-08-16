package model

import "github.com/jinzhu/gorm"

type Suite struct {
	gorm.Model
	Name    string `gorm:"size:300;not null;unique"`
	ThemeID uint
	Images  []Image
	IsRead  bool `gorm:"DEFAULT:false"`
	IsLike  bool `gorm:"DEFAULT:false"`
}

type SuitesQuery struct {
	IsLike bool `query:"is_like"`
}

type SuiteApiDownloadReq struct {
	Url string `query:"url"`
}
