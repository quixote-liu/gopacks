package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
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
	db := database.DB
	// auto migrate models
	if err := model.AutoMigrateModels(db); err != nil {
		log.Printf("automigrate models failed: %v", err)
		return
	}

	// *********************************************
	// *************** Create user *****************
	// *********************************************
	// create user
	if err := model.CreateUser(db, userCaseA); err != nil {
		log.Fatalf("create user failed: %v", err)
		return
	}

	// create users
	// if err := model.CreateUsers(db, []model.User{userCaseA, userCaseB}); err != nil {
	// 	log.Fatalf("create users failed: %v ", err)
	// 	return
	// }

	// *********************************************
	// *************** Find user *******************
	// *********************************************
	// first user.
	firstUser, err := model.FirstUser(db)
	if err != nil {
		log.Fatalf("find first user from databases failed: %v", err)
		return
	}
	fmt.Println("first user: ", *firstUser)

	// last user.
	lastUser, err := model.LastUser(db)
	if err != nil {
		log.Fatalf("find last user from the database failed: %v", err)
		return
	}
	fmt.Println("last user: ", *lastUser)

	// find user by name.
	user, err := model.FindUserByName(db, "lcs_2")
	if err != nil {
		log.Fatalf("find user from database by name %s failed: %v", "lcs_2", err)
		return
	}
	fmt.Printf("find user by name lcs_2: %v \n", *user)

	// find user by id
	user2, err := model.FindUserById(db, user.ID)
	if err != nil {
		log.Fatalf("find user by id %s failed: %v", user.ID, user)
		return
	}
	fmt.Printf("find user by id %s: %v\n", user.ID, *user2)

	// find all users;
	users, err := model.FindUsers(db)
	if err != nil {
		log.Fatalf("find users failed: %v", err)
		return
	}
	fmt.Println("all users: ", users)

	// *********************************************
	// *************** Update user *****************
	// *********************************************
	user.Name = "new_" + user.Name
	if err := model.UpdateUser(db, *user); err != nil {
		log.Fatalf("update user failed: %v", err)
		return
	}

	// find user by id
	user3, err := model.FindUserById(db, user.ID)
	if err != nil {
		log.Fatalf("find user by id %s failed: %v", user.ID, user)
		return
	}
	fmt.Printf("find user by id %s: %v \n", user.ID, *user3)

	// *********************************************
	// *************** Delete user *****************
	// *********************************************
	if err := model.DeleteUser(db, *user3); err != nil {
		log.Fatalf("delete user by id %s failed: %v", user3.ID, err)
		return
	}

	if _, err := model.FindUserById(db, user.ID); err == gorm.ErrRecordNotFound {
		fmt.Printf("the user which id is %s not found", user.ID)
	} else if err != nil {
		panic(err)
	}

	// GenSQLCase(database.DB)
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

func GenSQLCase(db *gorm.DB) {
	// Create SQL.
	res := model.Result{
		ID:   uuid.NewString(),
		Name: "wuhan",
		Age:  18,
	}
	createdSQL := model.CreateResultSQL(db, res)
	fmt.Println("createdSQL: ", createdSQL)

	// Find SQL.
	findSQL := model.FindResultSQL(db, res.ID)
	fmt.Println("findSQL: ", findSQL)
}
