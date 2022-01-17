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
	if err := db.AutoMigrate(&Student{}, &StudentCard{}); err != nil {
		log.Fatal(err)
		return
	}

	id := "63da95cf-839f-4c3c-ba75-6e059ee5a65a"
	studentCase := Student{
		ID:   id,
		Name: "xiaofang",
		StudentCard: StudentCard{
			Number: "20202111",
			Phone:  "10086",
		},
	}

	// *************************************
	// ************ Create *****************
	// *************************************
	studentCaseB := Student{
		ID:   "testing",
		Name: "xiaoming",
		StudentCard: StudentCard{
			Number: "20202001",
			Phone:  "10001",
		},
	}
	// if err := CreateStudent(db, studentCaseB); err != nil {
	// 	log.Fatalf("create student failed: %v", err)
	// 	return
	// }

	// *************************************
	// ************** Find *****************
	// *************************************
	studentA, err := FindStudentUsingGormPreload(db, studentCase.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("find student using gorm preload: %v\n", studentA)

	studentC, err := FindStudentByID(db, studentCase.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("find student by id %s: %v\n", studentCase.ID, studentC)

	studentCard, err := FindAssociatedStudentCardsByID(db, studentCase.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("associated studentCard: ", studentCard)

	// *************************************
	// ************** Update ***************
	// *************************************
	// studentCase.StudentCard.Phone = "10000"
	// if err := UpdateStudentCard(db, studentCase.StudentCard); err != nil {
	// 	log.Fatal(err)
	// 	return
	//
	// studentCaseB.Name = "xiao_ming_2"
	// studentCaseB.StudentCard.Phone = "10002"
	// if err := UpdateStudent(db, studentCaseB); err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// *************************************
	// ************** Delete ***************
	// *************************************
	if err := DeleteStudentByID(db, studentCaseB.ID, studentCaseB.StudentCard.Number); err != nil {
		log.Fatal(err)
		return
	}
	// if err := DeleteStudent(db, studentCaseB); err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
}

type Student struct {
	ID          string
	Name        string
	StudentCard StudentCard
}

type StudentCard struct {
	Number    string `gorm:"primaryKey"`
	StudentID string `gorm:"size:64"`
	Phone     string
}

// Create.
func CreateStudent(db *gorm.DB, s Student) error {
	return db.Create(&s).Error
}

// Find.
func FindStudentByID(db *gorm.DB, id string) (*Student, error) {
	s := &Student{}
	err := db.Where("id = ?", id).Find(s).Error
	return s, err
}

func FindStudentUsingGormPreload(db *gorm.DB, id string) (*Student, error) {
	s := &Student{}
	err := db.Preload("StudentCard").Where("id = ?", id).Find(s).Error
	return s, err
}

func FindAssociatedStudentCardsByID(db *gorm.DB, id string) (StudentCard, error) {
	studentCard := StudentCard{}
	err := db.Model(&Student{ID: id}).Association("StudentCard").Find(&studentCard)
	return studentCard, err
}

// Update.
func UpdateStudent(db *gorm.DB, s Student) error {
	return db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&s).Error
}

func UpdateStudentCard(db *gorm.DB, studentCard StudentCard) error {
	return db.Updates(&studentCard).Error
}

// Delete.
func DeleteStudentByID(db *gorm.DB, id, number string) error {
	return db.Select(clause.Associations).Delete(&Student{ID: id, StudentCard: StudentCard{
		Number: number,
	}}, "id = ?", id).Error
}

func DeleteStudent(db *gorm.DB, s Student) error {
	return db.Select(clause.Associations).Delete(&s).Error
}
