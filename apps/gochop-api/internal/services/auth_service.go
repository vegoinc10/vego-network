package services

import (
	"errors"

	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/models"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/repositories"
	"github.com/vegoinc10/vego-network/apps/gochop-api/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo   *repositories.UserRepository
	walletRepo *repositories.WalletRepository
}

func NewAuthService(
	userRepo *repositories.UserRepository,
	walletRepo *repositories.WalletRepository,
) *AuthService {

	return &AuthService{
		userRepo:   userRepo,
		walletRepo: walletRepo,
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

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	if user.Role == "vendor" {

		err = s.walletRepo.CreateWallet(user.ID)
		if err != nil {
			return nil, err
		}
	}

	user.PasswordHash = ""

	return user, nil
}
func (s *AuthService) Login(req *models.LoginRequest) (string, error) {

	user, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	)

	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
