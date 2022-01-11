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

	err := UserDemo(database.DB)
	if err != nil {
		log.Printf("user case error: %v", err)
		return
	}
}

var (
	userCaseA = model.User{
		Name:  "lcs_2",
		Email: "lcs@foxmail.com",
		Age:   18,
	}
	userCaseB = model.User{
		Name:  "lcs_3",
		Email: "lcs@email.com",
		Age:   18,
	}
)

func UserDemo(db *gorm.DB) error {
	// *********************************************
	// *************** Create user *****************
	// *********************************************

	// create user
	// if err := model.CreateUser(db, userCaseA); err != nil {
	// 	return fmt.Errorf("create user failed: %v", err)
	// }

	// create users
	// if err := model.CreateUsers(db, []model.User{userCaseA, userCaseB}); err != nil {
	// 	return fmt.Errorf("create users failed: %v ", err)
	// }

	// *********************************************
	// *************** Find user *******************
	// *********************************************

	// first user.
	firstUser, err := model.FirstUser(db)
	if err != nil {
		return fmt.Errorf("find first user from the mysql database failed: %v", err)
	}
	fmt.Println("first user: ", *firstUser)

	// last user.
	lastUser, err := model.LastUser(db)
	if err != nil {
		return fmt.Errorf("find last user from the mysql database failed: %v", err)
	}
	fmt.Println("last user: ", *lastUser)

	return nil
}
