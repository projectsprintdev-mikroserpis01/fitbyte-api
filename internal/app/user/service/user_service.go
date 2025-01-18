package service

import (
	"context"
	"database/sql"
	"errors"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/contracts"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/dto"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/pkg/bcrypt"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/pkg/jwt"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/pkg/log"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/pkg/validator"
)

type UserService interface {
	GetUserById(ctx context.Context, id int) (*dto.GetCurrentUserResponse, error)
	UpdateUserById(ctx context.Context, id int, req dto.UpdateUserRequest) (*dto.UpdateUserResponse, error)
}

type userService struct {
	repo      contracts.UserRepository
	jwt       jwt.JwtInterface
	bcrypt    bcrypt.BcryptInterface
	validator validator.ValidatorInterface
}

func NewUserService(
	repo contracts.UserRepository,
	jwt jwt.JwtInterface,
	bcrypt bcrypt.BcryptInterface,
	validator validator.ValidatorInterface,
) UserService {
	return &userService{
		repo:      repo,
		jwt:       jwt,
		bcrypt:    bcrypt,
		validator: validator,
	}
}

func (s *userService) GetUserById(ctx context.Context, id int) (*dto.GetCurrentUserResponse, error) {
	user, err := s.repo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	ret := dto.GetCurrentUserResponse{Email: user.Email, Name: user.Name, Preference: user.Preference, WeightUnit: user.WeightUnit, HeightUnit: user.HeightUnit, Weight: user.Weight, Height: user.Height, ImageURI: user.ImageURI}
	return &ret, nil
}

func (s *userService) UpdateUserById(ctx context.Context, id int, req dto.UpdateUserRequest) (*dto.UpdateUserResponse, error) {
	valErr := s.validator.Validate(req)
	if valErr != nil {
		return nil, valErr
	}

	if req.Email != nil {
		user, err := s.repo.GetUserByEmail(ctx, *req.Email)
		if err == nil { // found a user with the same email
			if user.ID != id {
				log.Info(log.LogInfo{
					"user id": user.ID,
					"id":      id,
				}, "[userService.UpdateUserById] id")

				return nil, domain.ErrUserEmailAlreadyExists
			}
		}

		if err != nil && !errors.Is(err, sql.ErrNoRows) { // some other error occurred
			return nil, err
		}
	}

	log.Info(log.LogInfo{
		"is req user image uri nil": req.UserImageUri == nil,
	}, "[userService.UpdateUserById] id")

	if req.UserImageUri != nil {
		u, err := url.ParseRequestURI(*req.UserImageUri)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusBadRequest, "invalid user image uri")
		}

		if u.Scheme == "" || u.Host == "" {
			return nil, fiber.NewError(fiber.StatusBadRequest, "invalid user image uri")
		}

		// Additional validation: Check if the host contains a domain or is not empty
		if !strings.Contains(u.Host, ".") {
			return nil, fiber.NewError(fiber.StatusBadRequest, "invalid company image uri")
		}
	}

	if req.CompanyImageUri != nil {
		u, err := url.ParseRequestURI(*req.CompanyImageUri)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusBadRequest, "invalid company image uri")
		}

		log.Info(log.LogInfo{
			"url": u,
		}, "[userService.UpdateUserById] company image uri")

		if u.Scheme == "" || u.Host == "" {
			return nil, fiber.NewError(fiber.StatusBadRequest, "invalid company image uri")
		}

		// Additional validation: Check if the host contains a domain or is not empty
		if !strings.Contains(u.Host, ".") {
			return nil, fiber.NewError(fiber.StatusBadRequest, "invalid company image uri")
		}
	}

	fields := []string{}
	args := []interface{}{}
	if req.Email != nil {
		fields = append(fields, "email")
		args = append(args, *req.Email)
	}
	if req.Name != nil {
		fields = append(fields, "name")
		args = append(args, *req.Name)
	}
	if req.UserImageUri != nil {
		fields = append(fields, "user_image_uri")
		args = append(args, *req.UserImageUri)
	}
	if req.CompanyName != nil {
		fields = append(fields, "company_name")
		args = append(args, *req.CompanyName)
	}
	if req.CompanyImageUri != nil {
		fields = append(fields, "company_image_uri")
		args = append(args, *req.CompanyImageUri)
	}

	_, err := s.repo.UpdateUserByIDSomeFields(ctx, id, fields, args)
	if err != nil {
		return nil, err
	}

	ret := dto.UpdateUserResponse{}

	return &ret, nil
}
