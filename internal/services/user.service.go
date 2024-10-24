package services

import (
	"context"

	"github.com/verlinof/softlancer-go/internal/database"
	"github.com/verlinof/softlancer-go/internal/models"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) GetUsers(ctx context.Context) (*[]models.User, error) {
	var users []models.User
	query := `
		SELECT *
		FROM users
	`
	err := database.DB.WithContext(ctx).Raw(query).Scan(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (u *UserService) GetUserbyEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	query := `
		SELECT *
		FROM users
		WHERE email = ?
	`

	err := database.DB.WithContext(ctx).Raw(query, email).Scan(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserService) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	err := database.DB.WithContext(ctx).Create(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) GetUserbyID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	err := database.DB.WithContext(ctx).Where("id = ?", id).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
