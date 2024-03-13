package service

import (
	"sesi_6/pkg/helpers"
	"sesi_6/pkg/models"

	"gorm.io/gorm"
)

type UserService struct {
	gorm *gorm.DB
}

func NewUserService(gorm *gorm.DB) *UserService {
	return &UserService{
		gorm: gorm,
	}
}

func (s *UserService) Register(request models.RegisterRequest) (models.User, error) {

	user := models.User{
		FullName: request.FullName,
		Username: request.Username,
		Email:    request.Email,
		Password: &request.Password,
	}

	err := s.gorm.Create(&user).Error

	return user, err
}

func (s *UserService) Login(request models.LoginRequest) (*models.LoginResponse, error) {
	var user models.User

	err := s.gorm.Where("username = ?", request.Username).First(&user).Error

	if err != nil {
		return nil, err
	}

	getPassword := *user.Password

	if !helpers.ComparePassword(getPassword, request.Password) {
		return nil, err
	}

	token, err := helpers.GenerateToken(user.ID, user.Username)

	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		AccessToken: token,
		User:        user,
	}, nil
}
