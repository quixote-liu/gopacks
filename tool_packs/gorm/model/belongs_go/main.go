package main

import (
	"fmt"
	"gopacks/tool_packs/gorm/database"
	"log"

	"gorm.io/gorm"
)

func main() {
	// init database
	if err := database.InitDB(); err != nil {
		log.Printf("init database failed: %v", err)
		return
	}
	db := database.DB

	// auto migrate models
	if err := db.AutoMigrate(&Member{}, &Company{}); err != nil {
		log.Fatal(err)
		return
	}

	// ************ Create *****************
	// programmerMember := model.Member{
	// 	ID:   uuid.NewString(),
	// 	Name: "programmer",
	// 	Company: model.Company{
	// 		ID:   uuid.NewString(),
	// 		Name: "tengxun",
	// 	},
	// }
	// err := model.CreateMember(db, programmerMember)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// ************ Create *****************
	id := "3849efc4-b6a9-4a1f-ad2e-88259619084a"
	memberA, err := FindMemberUsingPreLoad(db, id)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("memberA:", memberA)

	// Find member.
	memberB, err := FindMemberByID(db, id)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("memberB: ", memberB)
}

type Member struct {
	ID        string
	Name      string
	CompanyID string `gorm:"size:64"`
	Company   Company
}

type Company struct {
	ID   string
	Name string
}

func CreateMember(db *gorm.DB, m Member) error {
	return db.Create(&m).Error
}

func FindMemberByID(db *gorm.DB, id string) (*Member, error) {
	m := &Member{}
	err := db.Where("id = ?", id).Find(m).Error
	return m, err
}

func FindMemberUsingPreLoad(db *gorm.DB, id string) (*Member, error) {
	m := &Member{}
	err := db.Preload("Company").Where("id = ?", id).Find(m).Error
	return m, err
}

func FindMemberUsingJoins()
