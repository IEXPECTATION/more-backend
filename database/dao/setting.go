package dao

import "gorm.io/gorm"

type Setting struct {
	gorm.Model
	UserId uint
	BasicMode bool
}
