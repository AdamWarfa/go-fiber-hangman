package main

import (
	"go-fiber-hangman/initializers"
	"go-fiber-hangman/internal/controllers"
	"go-fiber-hangman/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})

	time.AfterFunc(24*time.Hour, func() {
		controllers.SetWord()
	})

	app := fiber.New()

	app.Get("/", controllers.GetWord)
	app.Post("/users", controllers.SaveUser)
	app.Put("/users/:id", controllers.UpdateUser)

	err := app.Listen(":6969")
	if err != nil {
		panic(err)
	}

}
