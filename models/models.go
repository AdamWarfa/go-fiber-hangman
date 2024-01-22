package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id      string `json:"id" gorm:"primaryKey"`
	Email   string `json:"email" gorm:"unique"`
	Streak  int    `json:"streak"`
	HiScore int    `json:"hiScore"`
}
