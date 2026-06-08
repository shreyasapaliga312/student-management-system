package models

import (
	"github.com/anaard/simple-student-management/pkg/config"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
)

var db *gorm.DB

// TODO: Ver uma forma de usar herança na class e student (métodos iguais)

func init() { // Function called before main
	if envLoadError := godotenv.Load(); envLoadError != nil { // Load environment variables
		log.Fatal("[ ERROR ] Failed to load .env file")
	}

	config.Connect()

	db = config.GetDB()

	db.AutoMigrate(&Student{}, &Class{}, &User{})
}
