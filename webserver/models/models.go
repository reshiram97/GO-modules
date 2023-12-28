package models

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type User struct {
	gorm.Model
	Id    uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name  string    `json:"name"`
	Password string `json:"password"`
	Email string    `json:"email"`
	Phone string    `json:"phone"`    
	Posts    []Post    `gorm:"foreignKey:UserID"`
}

type Post struct {
	gorm.Model
	Id    uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title string    `json:"title"`
	Info  string    `json:"info"`
	UserID uuid.UUID `json:"user_id" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID"`
	User  User      `json:"user"`
}
