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
	if err := db.AutoMigrate(&Member{}, &Company{}); err != nil {
		log.Fatal(err)
		return
	}

	// *************************************
	// ************ Create *****************
	// *************************************
	programmerMember := Member{
		ID:   "3849efc4-b6a9-4a1f-ad2e-88259619084a",
		Name: "programmer",
		Company: Company{
			ID:   "c7e8c7ff-c319-4b57-b28a-31584f1aeb5c",
			Name: "tengxun",
		},
	}
	err := CreateMember(db, programmerMember)
	if err != nil {
		log.Fatal(err)
		return
	}

	// *************************************
	// ************** Find *****************
	// *************************************
	memberA, err := FindMemberUsingGormPreLoad(db, programmerMember.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("find member using preload:", memberA)

	memberB, err := FindMemberByID(db, programmerMember.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("find member: ", memberB)

	company, err := FindAssociatedCompanyByID(db, programmerMember.Company.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Company: ", company)

	// *************************************
	// ************** Update ***************
	// *************************************
	programmerMember.Name = "xiaoming"
	programmerMember.Company.Name = "tengxun_2"
	if err := UpdateMember(db, programmerMember); err != nil {
		log.Fatal(err)
		return
	}

	// *************************************
	// ************** Delete ***************
	// *************************************
	if err := DeleteMemberByID(db, programmerMember); err != nil {
		log.Fatal(err)
		return
	}
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

// Create.
func CreateMember(db *gorm.DB, m Member) error {
	return db.Create(&m).Error
}

// Find.
func FindMemberByID(db *gorm.DB, id string) (*Member, error) {
	m := &Member{}
	err := db.Where("id = ?", id).Find(m).Error
	return m, err
}

func FindMemberUsingGormPreLoad(db *gorm.DB, id string) (*Member, error) {
	m := &Member{}
	err := db.Preload(clause.Associations).Where("id = ?", id).Find(m).Error
	return m, err
}

func FindAssociatedCompanyByID(db *gorm.DB, companyID string) (*Company, error) {
	c := &Company{}
	err := db.Model(&Member{CompanyID: companyID}).Association("Company").Find(c)
	return c, err
}

// Update
func UpdateMember(db *gorm.DB, m Member) error {
	return db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&m).Error
}

// Delete.
func DeleteMemberByID(db *gorm.DB, m Member) error {
	// return db.Select(clause.Associations).Delete(&m).Error
	return db.Delete(&m).Error
}
