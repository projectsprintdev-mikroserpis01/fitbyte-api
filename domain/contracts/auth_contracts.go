package contracts

import (
	"context"

	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/dto"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/entity"
)

type AuthRepository interface {
	EmailExists(ctx context.Context, email string) (bool, error)
	CreateUser(ctx context.Context, req dto.AuthRequest) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
}

type AuthService interface {
	Authenticate(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error)
}
