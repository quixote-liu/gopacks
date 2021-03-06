package model

import "gorm.io/gorm"

type Result struct {
	ID   string
	Name string
	Age  int
}

func CreateResultSQL(db *gorm.DB, res Result) string {
	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(&Result{}).Create(&res)
	})
	return sql
}

func FindResultSQL(db *gorm.DB, id string) string {
	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(&Result{}).Where("id = ?", id).Find(&Result{})
	})
	return sql
}
