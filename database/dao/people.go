package dao

import "gorm.io/gorm"

type People struct {
	gorm.Model
	Name string
}

func GetPeople(db *gorm.DB, name string) (p People) {
	db.First(&p, "name = ?", name)
	return p
}

func (p *People) SetPeopleName(db *gorm.DB, name string) {
}

func (p *People) SetPeoplePeopleId(db *gorm.DB, peopleid string) {
}
