package contracts

import (
	"context"

	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/dto"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/entity"
)

type UserRepository interface {
	GetUserById(ctx context.Context, id int) (*entity.User, error)
	UpdateUserById(ctx context.Context, id int, email string, name string, userImageUri string, companyName string, companyImageUri string) (int, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	UpdateUserByIDSomeFields(ctx context.Context, id int, fields []string, args []interface{}) (int, error)
}

type UserService interface {
	GetUserById(ctx context.Context, id int) (*dto.GetCurrentUserResponse, error)
	UpdateUserById(ctx context.Context, id int, req dto.UpdateUserRequest) (*dto.UpdateUserResponse, error)
}
