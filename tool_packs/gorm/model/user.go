package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Age       uint8
	Birthday  *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Hooks.
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

// Create user.
func CreateUser(db *gorm.DB, user User) error {
	return db.Create(&user).Error
}

func CreateUsers(db *gorm.DB, users []User) error {
	return db.Create(&users).Error
}

// Find user.
func FirstUser(db *gorm.DB) (*User, error) {
	u := &User{}
	err := db.First(u).Error
	return u, err
}

func TakeUser(db *gorm.DB) (*User, error) {
	u := &User{}
	err := db.Take(u).Error
	return u, err
}

func LastUser(db *gorm.DB) (*User, error) {
	u := &User{}
	err := db.Last(u).Error
	return u, err
}

func FindUserById(db *gorm.DB, id string) (*User, error) {
	u := &User{}
	err := db.Where("id = ?", id).First(u).Error
	return u, err
}

func FindUserByName(db *gorm.DB, name string) (*User, error) {
	u := &User{}
	err := db.Where("name = ?", name).First(u).Error
	return u, err
}

// Update user.
func UpdateUser(db *gorm.DB, user User) error {
	return db.Save(&user).Error
}

// Delete user.
func DeleteUser(db *gorm.DB, user User) error {
	return db.Delete(&user).Error
}
