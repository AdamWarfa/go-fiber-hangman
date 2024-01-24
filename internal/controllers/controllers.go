package controllers

import (
	"fmt"
	"go-fiber-hangman/initializers"
	"go-fiber-hangman/internal/services"
	"go-fiber-hangman/models"
	"net/http"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

var (
	randomWord     string = services.GetWord("https://random-word-api.herokuapp.com/word?number=1")
	randomWordLock sync.Mutex
)

func GetWord(c *fiber.Ctx) error {
	randomWordLock.Lock()
	defer randomWordLock.Unlock()

	fmt.Println(randomWord)

	c.Type("json")
	return c.Status(http.StatusOK).SendString(`["` + randomWord + `"]`)
}

func SetWord() {
	randomWordLock.Lock()
	defer randomWordLock.Unlock()

	// Update the random word
	randomWord = services.GetWord("https://random-word-api.herokuapp.com/word?number=1")

	// Reschedule the setWord function to be called again after 10 seconds
	time.AfterFunc(10*time.Hour, func() {
		SetWord()
	})
}

func SaveUser(c *fiber.Ctx) error {
	var body struct {
		Id      string `json:"id"`
		Email   string `json:"email"`
		Streak  int    `json:"streak"`
		HiScore int    `json:"hiScore"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	user := models.User{Id: body.Id, Email: body.Email, Streak: body.Streak, HiScore: body.HiScore}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save user"})
	}

	fmt.Printf("User %s saved in DB", user.Email)

	return c.Status(http.StatusOK).JSON(fiber.Map{"user": user})
}

func UpdateUser(c *fiber.Ctx) error {

	id := c.Params("id")

	var body struct {
		Id      string `json:"id"`
		Email   string `json:"email"`
		Streak  int    `json:"streak"`
		HiScore int    `json:"hiScore"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	user := models.User{Id: body.Id, Email: body.Email, Streak: body.Streak, HiScore: body.HiScore}

	result := initializers.DB.Model(&user).Where("id = ?", id).Updates(&user)

	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"user": user})
}
