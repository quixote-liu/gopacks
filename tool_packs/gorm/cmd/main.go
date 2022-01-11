package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"gopacks/tool_packs/gorm/database"
	"gopacks/tool_packs/gorm/model"
)

func main() {
	// init database
	if err := database.InitDB(); err != nil {
		log.Printf("init database failed: %v", err)
		return
	}
	// auto migrate models
	if err := model.AutoMigrateModels(database.DB); err != nil {
		log.Printf("automigrate models failed: %v", err)
		return
	}

	err := UserCase(database.DB)
	if err != nil {
		log.Printf("user case error: %v", err)
		return
	}
}

func UserCase(db *gorm.DB) error {
	userCaseA := model.User{
		Name:  "lcs_1",
		Email: "lcs_1@foxmail.com",
		Age:   18,
	}

	if err := model.CreateUser(db, userCaseA); err != nil {
		return fmt.Errorf("create user failed: %v", err)
	}

	return nil
}
