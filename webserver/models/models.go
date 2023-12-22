package models

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type User struct {
	gorm.Model
	Id    uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Phone string    `json:"phone"`    
}

type Post struct {
	gorm.Model
	Id     		uuid.UUID 		`json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title  		string    		`json:"title"`
	Info   		string    		`json:"info"`
	User 		uuid.UUID 		`gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID"`
}
