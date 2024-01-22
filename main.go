package main

import (
	"go-fiber-hangman/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", controllers.GetWord)
	app.Post("/users", controllers.SaveUser)
	app.Put("/users/:id", controllers.UpdateUser)

	err := app.Listen(":6969")
	if err != nil {
		panic(err)
	}

}
