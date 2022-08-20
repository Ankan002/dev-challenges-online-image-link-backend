package routes

import (
	"github.com/Ankan002/dev-challenges-online-image-link-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func UploadRouter(router fiber.Router) {
	router.Post("/upload", controllers.UploadFile)
}
