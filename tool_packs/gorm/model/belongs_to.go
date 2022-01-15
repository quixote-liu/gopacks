package model

import "gorm.io/gorm"

type Member struct {
	ID        string
	Name      string
	CompanyID string `gorm:"size:64"`
	Company   Company
}

type Company struct {
	ID   string
	Name string
}

func CreateMember(db *gorm.DB, m Member) error {
	return db.Create(&m).Error
}
