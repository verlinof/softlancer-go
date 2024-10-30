package pkg

import (
	"fmt"

	"github.com/verlinof/softlancer-go/config/app_config"
)

func PrefixBaseUrl(path string) *string {
	baseUrl := app_config.BASE_URL
	if path != "" {
		logoPath := fmt.Sprintf("%s%s", baseUrl, path)
		return &logoPath
	}
	return nil
}
