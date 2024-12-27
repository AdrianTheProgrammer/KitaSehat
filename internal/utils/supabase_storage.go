package utils

// import (
// 	"fmt"
// 	"os"
// 	"path/filepath"
// 	"strconv"

// 	"github.com/labstack/echo/v4"
// 	storage_go "github.com/supabase-community/storage-go"
// )

// func UploadAvatar(c echo.Context, userID int) (string, error, string) {
// 	// Connect to Supabase Storage
// 	supabase := storage_go.NewClient(os.Getenv("storage_endpoint"), os.Getenv("storage_key"), nil)

// 	// Bind Avatar File
// 	file, err := c.FormFile("avatar")
// 	if err != nil {
// 		return "", err, "error binding avatar file:"
// 	}

// 	// Open Avatar File
// 	src, err := file.Open()
// 	if err != nil {
// 		return "", err, "error opening avatar file"
// 	}
// 	defer src.Close()

// 	// Name the Avatar File by User ID
// 	path := fmt.Sprintf("avatar/%v.%v", strconv.Itoa(userID), filepath.Ext(file.Filename))

// 	// Upload to Supabase Storage
// 	response, err := supabase.UploadFile("KitaSehat", path, src)
// 	if err != nil {
// 		return "", err, response.Message
// 	}

// 	// Return Avatar URL
// 	return supabase.GetPublicUrl("KitaSehat", path).SignedURL, nil, ""
// }
