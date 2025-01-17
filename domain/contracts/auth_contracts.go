package contracts

import (
	"context"

	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/dto"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/entity"
)

type AuthRepository interface {
	EmailExists(ctx context.Context, email string) (bool, error)
	CreateUser(ctx context.Context, email string, password string) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
}

type AuthService interface {
	Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)
	Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error)
}
