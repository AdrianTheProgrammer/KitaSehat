package utils

import (
	"context"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo/v4"
)

type CloudinaryUtilityInterface interface {
	UploadAvatar(echo.Context, int) (string, error)
}

type CloudinaryUtility struct{}

func NewCloudinaryUtility() CloudinaryUtilityInterface {
	return &CloudinaryUtility{}
}

func (cu *CloudinaryUtility) UploadAvatar(c echo.Context, userID int) (string, error) {
	// Cloudinary Connection
	cloudinary, err := cloudinary.NewFromURL(os.Getenv("cloudinary"))
	if err != nil {
		return "", err
	}

	// Bind Avatar File
	file, err := c.FormFile("avatar")
	if err != nil {
		return "", err
	}

	// Open Avatar File
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Upload Avatar File to Cloudinary
	response, err := cloudinary.Upload.Upload(context.Background(), src, uploader.UploadParams{
		Folder:         "avatars",
		PublicID:       strconv.Itoa(userID),
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true),
	})
	if err != nil {
		return "", err
	}

	// Return Avatar URL
	return response.URL, nil
}
