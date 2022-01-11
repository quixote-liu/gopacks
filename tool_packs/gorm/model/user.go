package model

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"gopacks/tool_packs/gorm/database"
)

func init() {
	err := database.DB.AutoMigrate(&User{})
	if err != nil {
		log.Printf("migrate user model by gorm db failed: %v", err)
	}
}

type User struct {
	ID        string
	Name      string
	Email     string
	Age       uint8
	Birthday  *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	fmt.Println("before create user")
	return nil
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	_, err = uuid.Parse(u.ID)
	fmt.Println("after create user")
	return err
}

func CreateUser(db *gorm.DB, user User) error {
	return db.Create(&user).Error
}

func CreateUsers(db *gorm.DB, users []User) error {
	return db.Create(&users).Error
}
