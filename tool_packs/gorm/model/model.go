package model

import "gorm.io/gorm"

func AutoMigrateModels(db *gorm.DB) error {
	return db.AutoMigrate(&User{}, &Result{}, &Member{}, &Company{})
}
