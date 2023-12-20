package models

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)


type User struct {
	gorm.Model
	Id    uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

