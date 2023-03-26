package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nickname    string `json:"nickname" gorm:"unique"`
	Description string `json:"description"`
	Password    string `json:"-"`
}
