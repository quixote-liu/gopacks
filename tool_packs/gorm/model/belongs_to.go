package model

import "gorm.io/gorm"

type Member struct {
	ID        string
	Name      string
	CompanyID string
	Conpany   Company
}

type Company struct {
	ID   string
	Name string
}

func CreatePerson(db *gorm.DB, m Member) error {
	return db.Create(&m).Error
}
