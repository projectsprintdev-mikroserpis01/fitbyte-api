package service

import (
	"context"
	googleUuid "github.com/google/uuid"

	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/contracts"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/dto"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/pkg/bcrypt"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/pkg/jwt"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/pkg/uuid"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/pkg/validator"
)

type authService struct {
	authRepo  contracts.AuthRepository
	validator validator.ValidatorInterface
	uuid      uuid.UUIDInterface
	jwt       jwt.JwtInterface
	bcrypt    bcrypt.BcryptInterface
}

func NewAuthService(
	authRepo contracts.AuthRepository,
	validator validator.ValidatorInterface,
	uuid uuid.UUIDInterface,
	jwt jwt.JwtInterface,
	bcrypt bcrypt.BcryptInterface,
) contracts.AuthService {
	return &authService{
		authRepo:  authRepo,
		validator: validator,
		uuid:      uuid,
		jwt:       jwt,
		bcrypt:    bcrypt,
	}
}

func (s *authService) RegisterUser(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error) {
	// ..

	res := dto.RegisterResponse{
		ID: googleUuid.New(),
	}

	return res, nil
}

func (s *authService) LoginUser(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	// ..

	res := dto.LoginResponse{
		AccessToken: "accessToken",
	}

	return res, nil
}
