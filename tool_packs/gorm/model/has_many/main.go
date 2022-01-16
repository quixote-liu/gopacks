package main

import (
	"gopacks/tool_packs/gorm/database"
	"log"
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

}

type Client struct {
	ID    string
	Name  string
	Email []Email
}

type Email struct {
	ID       string
	Email    string
	ClientID string
}
