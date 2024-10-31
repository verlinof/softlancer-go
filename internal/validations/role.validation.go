package validations

import (
	"errors"
	"strings"

	"github.com/verlinof/softlancer-go/internal/database"
	"github.com/verlinof/softlancer-go/internal/models"
	"github.com/verlinof/softlancer-go/internal/requests"
)

func ValidateCreateRole(request *requests.RoleRequest) error {
	var validationErrors []string
	var roleExisted *models.Role

	//Validate Company Name
	if request.RoleName == "" {
		validationErrors = append(validationErrors, "Role name is required")
	} else {
		database.DB.First(&roleExisted, "role_name = ?", request.RoleName)
		if roleExisted.ID != "" {
			validationErrors = append(validationErrors, "Role name already exist")
		}
	}

	if len(validationErrors) > 0 {
		return errors.New(strings.Join(validationErrors, "; "))
	}

	return nil
}
