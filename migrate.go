package main

import (
	"go-fiber-hangman/initializers"
	"go-fiber-hangman/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func migrate() {
	initializers.DB.AutoMigrate(&models.User{})
}
