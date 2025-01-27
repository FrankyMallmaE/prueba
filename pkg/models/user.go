package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Age      string `json:"age"`
}
