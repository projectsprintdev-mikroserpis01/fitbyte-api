package service

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/contracts"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/dto"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/pkg/bcrypt"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/pkg/jwt"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/pkg/validator"
)

type authService struct {
	repo      contracts.AuthRepository
	validator validator.ValidatorInterface
	jwt       jwt.JwtInterface
	bcrypt    bcrypt.BcryptInterface
}

func NewAuthService(
	repo contracts.AuthRepository,
	validator validator.ValidatorInterface,
	jwt jwt.JwtInterface,
	bcrypt bcrypt.BcryptInterface,
) contracts.AuthService {
	return &authService{
		repo:      repo,
		validator: validator,
		jwt:       jwt,
		bcrypt:    bcrypt,
	}
}

func (s *authService) Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error) {
	valErr := s.validator.Validate(req)
	if valErr != nil {
		return dto.RegisterResponse{}, valErr
	}

	exists, err := s.repo.EmailExists(ctx, req.Email)
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	if exists {
		return dto.RegisterResponse{}, fiber.NewError(fiber.StatusConflict, "email already exists")
	}

	hashedPassword, err := s.bcrypt.Hash(req.Password)
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	req.Password = hashedPassword
	user, err := s.repo.CreateUser(ctx, req.Email, req.Password)
	if err != nil {
		fmt.Println("failed to CreateUser", err)

		return dto.RegisterResponse{}, err
	}

	token, err := s.jwt.Create(user.ID, user.Email)
	if err != nil {
		return dto.RegisterResponse{}, err
	}

	return dto.RegisterResponse{Email: req.Email, Token: token}, nil

}

func (s *authService) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	valErr := s.validator.Validate(req)
	if valErr != nil {
		return dto.LoginResponse{}, valErr
	}

	user, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return dto.LoginResponse{}, fiber.NewError(fiber.StatusNotFound, "user not found")
	}

	isValid := s.bcrypt.Compare(req.Password, user.Password)
	if !isValid {
		return dto.LoginResponse{}, domain.ErrCredentialsNotMatch
	}

	token, err := s.jwt.Create(user.ID, req.Email)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{Email: req.Email, Token: token}, nil

}
