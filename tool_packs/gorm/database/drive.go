package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	MysqlUserName = "testing"
	MysqlPassword = "testing"
	MysqlDBName   = "gorm"
	MysqlHost     = "127.0.0.1"
	MysqlPort     = "3306"
)

var defaultMysqlDriver = &mysqlDriver{
	port:     MysqlPort,
	host:     MysqlHost,
	userName: MysqlUserName,
	password: MysqlPassword,
	dbName:   MysqlDBName,
	gormConfig: gorm.Config{
		CreateBatchSize:      1000,
		FullSaveAssociations: true,
	},
}

type mysqlDriver struct {
	// mysql options
	port     string
	host     string
	userName string
	password string
	dbName   string

	// gorm config
	gormConfig gorm.Config
}

func (drv *mysqlDriver) dsn() string {
	u := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		drv.userName, drv.password, drv.host, drv.port, drv.dbName)
	return u
}

func (drv *mysqlDriver) connet() (*gorm.DB, error) {
	dsn := drv.dsn()
	return gorm.Open(mysql.Open(dsn), &drv.gormConfig)
}
