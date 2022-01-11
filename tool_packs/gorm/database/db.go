package database

import "gorm.io/gorm"

var DB *gorm.DB

func InitDB() error {
	var err error
	DB, err = defaultMysqlDriver.connetDB()
	if err != nil {
		return err
	}
	return nil
}
