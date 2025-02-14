package dao

import (
	"database/sql"

	"gorm.io/gorm"
)

type DreamItem struct {
	gorm.Model
	UserId      uint
	Name        string
	Cost        float64
	Description sql.NullString
	Completed   bool
}
