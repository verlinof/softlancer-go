package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/verlinof/softlancer-go/config/app_config"
	"github.com/verlinof/softlancer-go/validations"
)

func HandleUploadFile(c *gin.Context, form string, fileType []string, maxSize int) (string, error) {
	fileHeader, _ := c.FormFile(form)

	if fileHeader == nil {
		c.JSON(500, gin.H{
			"message": "File is required",
		})
		return "", fmt.Errorf("error upload file")
	}

	//Validation for file type
	if !validations.FileValidation(fileHeader, fileType) {
		c.JSON(500, gin.H{
			"message": "File type is not allowed",
		})
		return "", fmt.Errorf("file type is not allowed")
	}

	// // Validate file size
	// if fileHeader.Size > int64(maxSize) {
	// 	return "", fmt.Errorf("file size exceeds limit")
	// }

	//Akan mengembalikan path file
	extensionFile := filepath.Ext(fileHeader.Filename)
	filename := uuid.New().String() + extensionFile
	errUpload := c.SaveUploadedFile(fileHeader, fmt.Sprintf(app_config.STATIC_DIR+"/%s", filename))
	if errUpload != nil {
		c.JSON(500, gin.H{
			"message": errUpload.Error(),
		})
		return "", fmt.Errorf("error upload file")
	}

	// return filename
	pathFile := fmt.Sprintf(app_config.STATIC_PATH+"/%s", filename)

	return pathFile, nil
}

func HandleRemoveFile(filename string) error {
	// filename := c.Param("filename")
	errRemove := os.Remove(fmt.Sprintf(".%s", filename))
	if errRemove != nil {
		return fmt.Errorf("error remove file")
	}
	return nil
}
