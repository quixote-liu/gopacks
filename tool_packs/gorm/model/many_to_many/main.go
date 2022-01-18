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
	if err := db.AutoMigrate(&Person{}, &Language{}); err != nil {
		log.Fatal(err)
		return
	}

	languageAID := "ad58e657-6ca9-45ee-9ca7-c49111fc91ae"
	languageBID := "242971d9-ad52-4cd9-af7b-c729551a489e"
	personCase := Person{
		ID:   "5be970ce-d3c0-46c9-aab0-6cf9b715b501",
		Name: "xiaoming",
		Age:  19,
		Languages: []Language{
			{ID: languageAID, Name: "chinese"},
			{ID: languageBID, Name: "English"},
		},
	}

	// *************************************
	// *************** Create **************
	// *************************************
	// if err := CreatePerson(db, personCase); err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// *************************************
	// *************** Find ****************
	// *************************************
	p, err := FindPersonByID(db, personCase.ID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("find person by id(%s) : %v\n", personCase.ID, p)

	l, err := FindLanguageByID(db, languageAID)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("find language by id(%s) : %v\n", languageAID, l)

	// *************************************
	// *************** Update **************
	// *************************************
	// personCase.Name = "xiaoming"
	// for i, lang := range personCase.Languages {
	// 	personCase.Languages[i].Name = strings.ToUpper(lang.Name)
	// }
	// if err := UpdatePerson(db, personCase); err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// *************************************
	// *************** Delete **************
	// *************************************
	// if err := DeletePerson(db, personCase.ID); err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
}

type Person struct {
	ID        string `gorm:"size:64"`
	Name      string
	Age       int
	Languages []Language `gorm:"many2many:person_languages"`
}

type Language struct {
	ID      string `gorm:"size:64"`
	Name    string
	Persons []Person `gorm:"many2many:person_languages"`
}

// Create.
func CreatePerson(db *gorm.DB, p Person) error {
	return db.Create(&p).Error
}

// Update.
func UpdatePerson(db *gorm.DB, p Person) error {
	return db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&p).Error
}

// Find.
func FindPersonByID(db *gorm.DB, id string) (*Person, error) {
	p := &Person{}
	err := db.Preload(clause.Associations).Find(p).Error
	return p, err
}

func FindLanguageByID(db *gorm.DB, id string) (*Language, error) {
	l := &Language{}
	err := db.Preload(clause.Associations).Find(l).Error
	return l, err
}

// Delete.
func DeletePerson(db *gorm.DB, id string) error {
	return db.Select(clause.Associations).Delete(&Person{ID: id}).Error
}

// SQL.
func AutomirageSQL(db *gorm.DB) string {
	sql := db.ToSQL(func(db *gorm.DB) *gorm.DB {
		if err := db.AutoMigrate(&Person{}, &Language{}); err != nil {
			log.Fatal(err)
			return nil
		}
		return db
	})
	return sql
}
