package controllers

import (
	"context"
	"github.com/Ankan002/dev-challenges-online-image-link-backend/config"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"strings"
	"time"
)

func UploadFile(ctx *fiber.Ctx) error {

	photoFile, err := ctx.FormFile("photo")

	fileExtension := strings.Split(photoFile.Filename, ".")[len(strings.Split(photoFile.Filename, "."))-1]

	if fileExtension != "jpeg" && fileExtension != "png" && fileExtension != "jpg" && fileExtension != "gif" && fileExtension != "web" && fileExtension != "svg" && fileExtension != "ico" {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid image format",
		})
	}

	if photoFile.Size > (1024 * 1024 * 20) {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "File size should be more than 20 MB",
		})
	}

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "File not received",
		})
	}

	photo, fileOpenError := photoFile.Open()

	if fileOpenError != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "File cannot be opened",
		})
	}

	cloudinaryInstance := config.GetCloudinaryInstance()

	asyncContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	response, uploadError := cloudinaryInstance.Upload.Upload(asyncContext, photo, uploader.UploadParams{PublicID: "online-image-link/" + uuid.New().String()})

	log.Println(photo.Close())

	if uploadError != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   uploadError.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"url": response.SecureURL,
		},
	})
}
