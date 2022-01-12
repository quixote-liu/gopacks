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
	if err := model.CreateUser(db, userCaseA); err != nil {
		return fmt.Errorf("create user failed: %v", err)
	}

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

	// find user by name.
	user, err := model.FindUserByName(db, "lcs_2")
	if err != nil {
		return fmt.Errorf("find user by name %s failed: %v", "lcs_2", err)
	}
	fmt.Printf("find user by name lcs_2: %v \n", *user)

	// find user by id
	user2, err := model.FindUserById(db, user.ID)
	if err != nil {
		return fmt.Errorf("find user by id %s failed: %v", user.ID, user)
	}
	fmt.Printf("find user by id %s: %v\n", user.ID, *user2)

	// *********************************************
	// *************** Update user *****************
	// *********************************************
	user.Name = "new_" + user.Name
	if err := model.UpdateUser(db, *user); err != nil {
		return fmt.Errorf("update user failed: %v", err)
	}

	// find user by id
	user3, err := model.FindUserById(db, user.ID)
	if err != nil {
		return fmt.Errorf("find user by id %s failed: %v", user.ID, user)
	}
	fmt.Printf("find user by id %s: %v \n", user.ID, *user3)

	// *********************************************
	// *************** Delete user *****************
	// *********************************************
	// delete user by id.
	if err := model.DeleteUser(db, *user3); err != nil {
		return fmt.Errorf("delete user by id %s failed: %v", user3.ID, err)
	}

	if _, err := model.FindUserById(db, user.ID); err == gorm.ErrRecordNotFound {
		fmt.Printf("the user which id is %s not found", user.ID)
	} else if err != nil {
		panic(err)
	}

	return nil
}
