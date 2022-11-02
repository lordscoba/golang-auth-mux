package models

import(
	"github.com/jinzhu/gorm"
)


type User struct{
	gorm.Model
	Name string `gorm:""json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Token	string	`json:"token"`
	User_type string `json:"user_type"`
	Description string `json:"description"`
}

// var Db *gorm.DB