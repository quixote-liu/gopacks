package main

import (
	"fmt"
	"gopacks/tool_packs/gorm/database"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {
	// init database
	if err := database.InitDB(); err != nil {
		log.Printf("init database failed: %v", err)
		return
	}

	db := database.DB

	// auto migrate models
	if err := db.AutoMigrate(&Client{}, &Email{}); err != nil {
		log.Fatal(err)
		return
	}

	clientCase := Client{
		ID:   "4ed6df5d-0721-42b7-81ae-322d76d2f496",
		Name: "xiaofang",
		Emails: []Email{
			{ID: "7e425b8a-1a6d-4029-84e3-7acdcdb044e9", Email: "222@foxmail.com"},
			{ID: "1e94e529-0578-4ae1-b78c-a954d074ce69", Email: "333@foxmail.com"},
		},
	}

	// *************************************
	// *************** Create **************
	// *************************************
	// if err := CreateClient(db, clientCase); err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// *************************************
	// *************** Find ****************
	// *************************************
	c, err := FindClientByID(db, clientCase.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("find client by ID [%s] : %v\n", clientCase.ID, *c)

	// *************************************
	// *************** Update **************
	// *************************************
	// clientCase.Name = "xiaoming"
	// for i, e := range clientCase.Emails {
	// 	clientCase.Emails[i].Email = "new_" + e.Email
	// }
	// if err := UpdateClient(db, clientCase); err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// *************************************
	// *************** Delete **************
	// *************************************
	// if err := DeleteClientByID(db, clientCase.ID); err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
}

type Client struct {
	ID     string
	Name   string
	Emails []Email
}

type Email struct {
	ID       string
	Email    string
	ClientID string `gorm:"size:64"`
}

// Create.
func CreateClient(db *gorm.DB, c Client) error {
	return db.Create(&c).Error
}

// Find.
func FindClientByID(db *gorm.DB, id string) (*Client, error) {
	c := &Client{}
	err := db.Preload(clause.Associations).Find(c).Error
	return c, err
}

// Update
func UpdateClient(db *gorm.DB, c Client) error {
	return db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&c).Error
}

// Delete.
func DeleteClientByID(db *gorm.DB, id string) error {
	return db.Select(clause.Associations).Delete(&Client{ID: id}).Error
}
