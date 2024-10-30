package pkg

import (
	"fmt"

	"github.com/verlinof/softlancer-go/config/app_config"
)

// PrefixBaseUrl takes a path and prepends the app's base URL to it.
// If the path is empty, it returns nil.
// This is used to generate full URLs for images, etc. in the API responses.
func PrefixBaseUrl(path string) *string {
	baseUrl := app_config.BASE_URL
	if path != "" {
		logoPath := fmt.Sprintf("%s%s", baseUrl, path)
		return &logoPath
	}
	return nil
}
