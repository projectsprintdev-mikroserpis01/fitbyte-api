package repository

import (
	"context"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/contracts"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/dto"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/entity"
)

type userRepository struct {
	conn *sqlx.DB
	tx   *sqlx.Tx
}

func NewUserRepository(conn *sqlx.DB) contracts.UserRepository {
	return &userRepository{
		conn: conn,
	}
}

func (u *userRepository) BeginTransaction(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) CommitTransaction() error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) RollbackTransaction() error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) GetUsers(ctx context.Context, query dto.GetUsersQuery) ([]entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) GetUserByField(ctx context.Context, field, value string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) CreateUser(ctx context.Context, user *entity.User) (uuid.UUID, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) UpdateUser(ctx context.Context, user *entity.User) (uuid.UUID, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) SoftDeleteUser(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) DeleteUser(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) RestoreUser(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) CountUsers(ctx context.Context, query dto.GetUsersStatsQuery) (int64, error) {
	//TODO implement me
	panic("implement me")
}
