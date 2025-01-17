package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/contracts"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/entity"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) contracts.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) EmailExists(ctx context.Context, email string) (bool, error) {
	var exists bool
	err := r.db.Get(&exists, "SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)", email)
	return exists, err
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE email=$1", email)
	return user, err
}

func (r *userRepository) GetUserById(ctx context.Context, id int) (*entity.User, error) {
	var user entity.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE id=$1", id)
	return &user, err
}

func (r *userRepository) UpdateUserById(ctx context.Context, id int, email string, name string, userImageUri string, companyName string, companyImageUri string) (int, error) {

	_, err := r.GetUserByEmail(ctx, email)
	if err == nil { // successfully found a user with the same email
		return 0, domain.ErrUserEmailAlreadyExists
	}
	result, err := r.db.ExecContext(ctx, "UPDATE users SET name = $1, user_image_uri = $2, company_name = $3, company_image_uri = $4 WHERE id = $5", name, userImageUri, companyName, companyImageUri, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, _ := result.RowsAffected()
	return int(rowsAffected), nil
}

func (r *userRepository) UpdateUserByIDSomeFields(ctx context.Context, id int, fields []string, args []interface{}) (int, error) {
	query := `UPDATE users SET `
	for i, field := range fields {
		query += fmt.Sprintf("%s = $%d", field, i+1)
		if i != len(fields)-1 {
			query += ", "
		}
	}
	query += fmt.Sprintf(" WHERE id = $%d", len(fields)+1)

	args = append(args, id)

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	rowsAffected, _ := result.RowsAffected()
	return int(rowsAffected), nil
}
