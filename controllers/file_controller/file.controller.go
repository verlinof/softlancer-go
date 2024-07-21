package file_controller

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/verlinof/softlancer-go/config/app_config"
	"github.com/verlinof/softlancer-go/utils"
)

func HandleUploadFile(c *gin.Context) {
	fileHeader, _ :=	c.FormFile("file")

	if(fileHeader == nil) {
		c.JSON(500, gin.H{
			"message": "File is required",
		})
		return
	}

	//Validation for file type
	if(!utils.FileValidation(fileHeader, []string{"image/png", "image/jpeg"})) {
		c.JSON(500, gin.H{
			"message": "File type is not allowed",
		})
		return
	}

	//Akan mengembalikan path file
	extensionFile := filepath.Ext(fileHeader.Filename)
	filename := uuid.New().String() + extensionFile
	errUpload := c.SaveUploadedFile(fileHeader, fmt.Sprintf(app_config.STATIC_DIR + "/%s", filename))
	if(errUpload != nil) {
		c.JSON(500, gin.H{
			"message": errUpload.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success Upload File",
		"path": filename,
	})
}

func HandleRemoveFile(c *gin.Context) {
	filename := c.Param("filename")

	errRemove := os.Remove(fmt.Sprintf(app_config.STATIC_DIR + "/%s", filename))
	if(errRemove != nil) {
		c.JSON(500, gin.H{
			"message": errRemove.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success Delete File",
		"path": filename,
	})
}