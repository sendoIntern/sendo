package utils

import (
	"context"
	"log"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/joho/godotenv"
)

func UploadToCloudinary(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUD_NAME"),
		os.Getenv("CLOUD_API_KEY"),
		os.Getenv("CLOUD_API_SECRET"),
	)
	if err != nil {
		log.Println("Cloudinary config error:", err)
		return "", err
	}

	uploadResult, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{
		PublicID: fileHeader.Filename,
	})
	if err != nil {
		log.Println("Upload error:", err)
		return "", err
	}

	return uploadResult.SecureURL, nil
}
