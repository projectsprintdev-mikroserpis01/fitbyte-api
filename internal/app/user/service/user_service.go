package service

import (
	"context"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
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

	log.Info(log.LogInfo{
		"is req user image uri nil": req.ImageURI == nil,
	}, "[userService.UpdateUserById] id")

	if req.ImageURI != nil {
		u, err := url.ParseRequestURI(*req.ImageURI)
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

	fields := []string{}
	args := []interface{}{}

	if req.Name != nil {
		fields = append(fields, "name")
		args = append(args, *req.Name)
	}
	if req.ImageURI != nil {
		fields = append(fields, "image_uri")
		args = append(args, *req.ImageURI)
	}

	fields = append(fields, "preference")
	args = append(args, req.Preference)

	fields = append(fields, "weight_unit")
	args = append(args, req.WeightUnit)

	fields = append(fields, "height_unit")
	args = append(args, req.HeightUnit)

	fields = append(fields, "height")
	args = append(args, req.Height)

	fields = append(fields, "weight")
	args = append(args, req.Weight)

	_, err := s.repo.UpdateUserByIDSomeFields(ctx, id, fields, args)
	if err != nil {
		return nil, err
	}

	ret := dto.UpdateUserResponse{}

	return &ret, nil
}
