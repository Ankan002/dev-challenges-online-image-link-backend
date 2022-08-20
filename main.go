package main

import (
	"github.com/Ankan002/dev-challenges-online-image-link-backend/config"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func main() {
	if os.Getenv("GO_ENV") != "production" {
		config.LoadEnv()
	}

	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		log.Println(os.Getenv("APP_NAME"))
		return ctx.Status(200).JSON(fiber.Map{
			"success": true,
			"message": "Welcome to the Online Image Link API",
		})
	})

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
