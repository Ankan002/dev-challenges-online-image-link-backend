package main

import (
	"github.com/Ankan002/dev-challenges-online-image-link-backend/config"
	"github.com/Ankan002/dev-challenges-online-image-link-backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"os"
)

func main() {
	if os.Getenv("GO_ENV") != "production" {
		config.LoadEnv()
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

	router := app.Group("/api")

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(fiber.Map{
			"success": true,
			"message": "Welcome to the Online Image Link API",
		})
	})

	routes.UploadRouter(router)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
