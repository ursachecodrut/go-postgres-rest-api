package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	name  string
	email string
	age   uint8
}
