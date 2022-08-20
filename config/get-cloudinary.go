package config

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"log"
	"os"
)

var cloudinaryInstance *cloudinary.Cloudinary

func getCloudinaryInstance() *cloudinary.Cloudinary {
	if cloudinaryInstance != nil {
		return cloudinaryInstance
	}

	cloudinaryCloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	cloudinaryApiKey := os.Getenv("CLOUDINARY_API_KEY")
	cloudinaryApiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	cld, err := cloudinary.NewFromParams(cloudinaryCloudName, cloudinaryApiKey, cloudinaryApiSecret)
	
	if err != nil {
		log.Fatal(err)
	}

	cloudinaryInstance = cld

	return cloudinaryInstance
}
