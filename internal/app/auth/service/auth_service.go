package service

import (
	"context"
	"errors"

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

func (s *authService) Authenticate(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error) {
	valErr := s.validator.Validate(req)
	if valErr != nil {
		return dto.AuthResponse{}, valErr
	}

	switch req.Action {
	case "create":
		exists, err := s.repo.EmailExists(ctx, req.Email)
		if err != nil {
			return dto.AuthResponse{}, err
		}

		if exists {
			return dto.AuthResponse{}, fiber.NewError(fiber.StatusConflict, "email already exists")
		}

		hashedPassword, err := s.bcrypt.Hash(req.Password)
		if err != nil {
			return dto.AuthResponse{}, err
		}

		req.Password = hashedPassword
		user, err := s.repo.CreateUser(ctx, req)
		if err != nil {
			return dto.AuthResponse{}, err
		}

		token, err := s.jwt.Create(user.ID, user.Email)
		if err != nil {
			return dto.AuthResponse{}, err
		}

		return dto.AuthResponse{Email: req.Email, Token: token}, nil

	case "login":
		user, err := s.repo.GetUserByEmail(ctx, req.Email)
		if err != nil {
			return dto.AuthResponse{}, fiber.NewError(fiber.StatusNotFound, "user not found")
		}

		isValid := s.bcrypt.Compare(req.Password, user.Password)
		if !isValid {
			return dto.AuthResponse{}, domain.ErrCredentialsNotMatch
		}

		token, err := s.jwt.Create(user.ID, req.Email)
		if err != nil {
			return dto.AuthResponse{}, err
		}

		return dto.AuthResponse{Email: req.Email, Token: token}, nil

	default:
		return dto.AuthResponse{}, errors.New("invalid action")
	}

}
