package validations

import (
	"github.com/verlinof/softlancer-go/database"
	"github.com/verlinof/softlancer-go/models"
	"github.com/verlinof/softlancer-go/requests"
)

func ValidateCreateRole(request *requests.RoleRequest) []string {
	var validationErrors []string
	var roleExisted *models.Role

	//Validate Company Name
	if request.RoleName == "" {
		validationErrors = append(validationErrors, "Role name is required")
	} else {
		database.DB.First(&roleExisted, "role_name = ?", request.RoleName)
		if roleExisted.ID != nil {
			validationErrors = append(validationErrors, "Role name already exist")
		}
	}

	return validationErrors
}
