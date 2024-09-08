package services

import (
	"context"

	"github.com/verlinof/softlancer-go/internal/database"
	"github.com/verlinof/softlancer-go/internal/responses"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) GetUsers(ctx context.Context) ([]responses.UserResponse, error) {
	var userRes []responses.UserResponse
	err := database.DB.WithContext(ctx).Table("users").
		Select("id, email, name, address").
		Scan(&userRes).Error

	if err != nil {
		return nil, err
	}
	return userRes, nil
}

// func (u *UserService) CreateUser(ctx context.Context, user *models.User) (*responses.UserResponse, error) {
// 	err := database.DB.WithContext(ctx).Create(&user).Error

// 	if err != nil {
// 		return nil, err
// 	}

// }
