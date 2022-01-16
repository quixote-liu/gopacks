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
		},
	}

	// *************************************
	// ************ Create *****************
	// *************************************
	// if err := CreateStudent(db, studentCase); err != nil {
	// 	log.Fatalf("create student failed: %v", err)
	// 	return
	// }

	studentCaseB := Student{
		ID:   "05147eb7-4d73-47c2-971b-298f061b4feb",
		Name: "xiaoming",
		StudentCard: StudentCard{
			Number: "20202112",
		},
	}
	// if err := CreateStudent(db, studentCaseB); err != nil {
	// 	log.Fatalf("create student failed: %v", err)
	// 	return
	// }

	// *************************************
	// ************** Find *****************
	// *************************************
	studentA, err := FindStudentUsingGormPreload(db, studentCaseB.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("find student using gorm preload: %v\n", studentA)

	studentB, err := FindStudentUsingGormJoins(db, studentCaseB.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("find student using gorm joins: %v\n", studentB)

	studentC, err := FindStudentByID(db, studentCase.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("find student by id %s: %v\n", studentCase.ID, studentC)

	// *************************************
	// ************** Update ***************
	// *************************************
	// studentCaseB.Name = "xiao_ming_13"
	// studentCaseB.StudentCard.Number = "20202013"
	// if err := UpdateStudent(db, studentCaseB); err != nil {
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

func FindStudentUsingGormJoins(db *gorm.DB, id string) (*Student, error) {
	s := &Student{}
	err := db.Joins("StudentCard").Where("students.id = ?", id).Find(s).Error
	return s, err
}

// Update.
func UpdateStudent(db *gorm.DB, s Student) error {
	return db.Preload("StudentCard").Updates(&s).Error
}
