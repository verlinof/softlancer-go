package utils

import "mime/multipart"

func FileValidation(fileHeader *multipart.FileHeader, fileType []string) bool {
	contentType := fileHeader.Header.Get("Content-Type")
	result := false

	for _, typefile := range fileType {
		if(typefile == contentType) {
			result = true
			break
		}
	}

	return result
}