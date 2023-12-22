package database

import (
	"log"
	"os"
	"github.com/reshiram97/webserver/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %s", err)
    }
    username := os.Getenv("USER_NAME")
    password := os.Getenv("PASSWORD")
    database := os.Getenv("DATABASE")
    dbURL := "postgres://" + username + ":" + password + "@localhost:5432/" + database + "?sslmode=disable" // Added sslmode=disable
    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
    if err != nil {
        log.Fatalln(err)
    }

    // AutoMigrate with foreign key constraints
    db.AutoMigrate(&models.User{}, &models.Post{})

    return db
}
