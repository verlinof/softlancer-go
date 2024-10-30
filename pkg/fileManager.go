package pkg

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/verlinof/softlancer-go/config/app_config"
	"github.com/verlinof/softlancer-go/internal/validations"
)

// HandleUploadFile upload a file to static directory and return the path file
// if file is not provided, it will return error
// if file type is not allowed, it will return error
// if file size exceeds limit, it will return error
//
// Parameters:
// c *gin.Context - context of gin framework
// folderName string - folder name where the file will be stored
// form string - form name in the request
// fileType []string - allowed file types
// maxSize int - max size of the file in bytes
//
// Returns:
// string - path of the uploaded file
// error - error if something went wrong
func HandleUploadFile(c *gin.Context, folderName string, form string, fileType []string, maxSize int) (string, error) {
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
	errUpload := c.SaveUploadedFile(fileHeader, fmt.Sprintf(app_config.STATIC_DIR+"/%s/%s", folderName, filename))
	if errUpload != nil {
		c.JSON(500, gin.H{
			"message": errUpload.Error(),
		})
		return "", fmt.Errorf("error upload file")
	}

	// return filename
	pathFile := fmt.Sprintf(app_config.STATIC_PATH+"/%s/%s", folderName, filename)

	return pathFile, nil
}

// HandleRemoveFile removes a file from static directory
//
// Parameters:
// filename string - name of the file
//
// Returns:
// error - error if something went wrong
func HandleRemoveFile(filename string) error {
	// filename := c.Param("filename")
	errRemove := os.Remove(fmt.Sprintf(".%s", filename))
	if errRemove != nil {
		return fmt.Errorf("error remove file")
	}
	return nil
}
