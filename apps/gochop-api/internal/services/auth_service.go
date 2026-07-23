package services

import (
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/repositories"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/utils"
)

type AuthService struct {
	UserRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{
		UserRepo: userRepo,
	}
}

func (s *AuthService) Register(req *models.RegisterRequest) (*models.User, error) {

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		FullName:     req.FullName,
		Email:        req.Email,
		Phone:        req.Phone,
		PasswordHash: hashedPassword,
		Role:         req.Role,
		IsVerified:   false,
		IsActive:     true,
	}

	err = s.UserRepo.Create(user)
	if err != nil {
		return nil, err
	}

	user.PasswordHash = ""

	return user, nil
}
