package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/contracts"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/entity"
)

type authRepository struct {
	conn *sqlx.DB
	tx   *sqlx.Tx
}

func NewAuthRepository(conn *sqlx.DB) contracts.AuthRepository {
	return &authRepository{
		conn: conn,
	}
}

func (a *authRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User

	// ..

	return &user, nil
}

func (a *authRepository) RegisterUser(ctx context.Context, user entity.User) (uuid.UUID, error) {
	// ..

	return user.ID, nil
}
