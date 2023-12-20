package database

import (
	"log"
	"os"
	"webserver/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
);

func Init() *gorm.DB {
    err := godotenv.Load(".env")
    if err != nil{
        log.Fatalf("Error loading .env file: %s", err)
    }
    username:= os.Getenv("USER_NAME")
    password:= os.Getenv("PASSWORD")
    database:= os.Getenv("DATABASE")
	dbURL := "postgres://"+username+":"+password+"@localhost:5432/"+database;
    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
    if err != nil {
        log.Fatalln(err)
    }
    db.AutoMigrate(&models.User{})
    return db
}