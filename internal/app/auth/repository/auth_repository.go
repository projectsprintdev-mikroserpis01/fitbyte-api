package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/contracts"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/dto"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/entity"
)

type authRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) contracts.AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) EmailExists(ctx context.Context, email string) (bool, error) {
	var exists bool
	err := r.db.GetContext(ctx, &exists, "SELECT EXISTS(SELECT 1 FROM managers WHERE email=$1)", email)
	return exists, err
}

func (r *authRepository) CreateUser(ctx context.Context, req dto.AuthRequest) (entity.User, error) {
	_, err := r.db.Exec("INSERT INTO managers (email, password) VALUES ($1, $2)", req.Email, req.Password)
	if err != nil {
		return entity.User{}, err
	}

	return r.GetUserByEmail(ctx, req.Email)
}

func (r *authRepository) UpdateUser(ctx context.Context, req dto.UserProfile) (entity.User, error) {
	_, err := r.db.ExecContext(ctx, "UPDATE managers SET name = $1, user_image_uri = $2, company_name = $3, company_image_uri = $4 WHERE email = $5", req.Name, req.UserImageUri, req.CompanyName, req.CompanyImageUri, req.Email)
	if err != nil {
		return entity.User{}, err
	}

	return r.GetUserByEmail(ctx, req.Email)
}

func (r *authRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	err := r.db.GetContext(ctx, &user, "SELECT * FROM managers WHERE email=$1", email)
	return user, err
}
