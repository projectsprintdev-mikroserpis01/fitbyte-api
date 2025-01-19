package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/contracts"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/dto"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/entity"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/enums"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/pkg/log"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/pkg/uuid"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/pkg/validator"
)

type activityStruct struct {
	activityRepo contracts.ActivityRepository
	validator    validator.ValidatorInterface
	uuid         uuid.UUIDInterface
}

func NewActivityService(activityRepo contracts.ActivityRepository, validator validator.ValidatorInterface, uuid uuid.UUIDInterface) contracts.ActivityService {
	return &activityStruct{
		activityRepo: activityRepo,
		validator:    validator,
		uuid:         uuid,
	}
}

// CreateActivity implements contracts.ActivityService.
func (a *activityStruct) CreateActivity(ctx context.Context, req dto.CreateActivityRequest) (dto.CreateActivityResponse, error) {
	log.Info(log.LogInfo{
		"req": req,
	}, "[CreateActivity] Request")

	valErr := a.validator.Validate(req)
	if valErr != nil {
		return dto.CreateActivityResponse{}, valErr
	}

	id, err := a.uuid.NewV7()
	if err != nil {
		return dto.CreateActivityResponse{}, err
	}

	activity := entity.Activity{
		ID:                id,
		ActivityType:      enums.ActivityTypesReverse[req.ActivityType],
		DoneAt:            req.DoneAt,
		DurationInMinutes: req.DurationInMinutes,
		UserID:            req.UserID,
		CaloriesBurned:    enums.Calories[req.ActivityType] * req.DurationInMinutes,
	}

	err = a.activityRepo.CreateActivity(ctx, activity)
	if err != nil {
		return dto.CreateActivityResponse{}, err
	}
	res := dto.CreateActivityResponse{
		ActivityID:        activity.ID,
		AcitivityType:     req.ActivityType,
		DoneAt:            req.DoneAt,
		DurationInMinutes: req.DurationInMinutes,
		CaloriesBurned:    activity.CaloriesBurned,
		CreatedAt:         activity.CreatedAt,
		UpdateAt:          time.Now(),
	}

	return res, nil
}

func (a *activityStruct) GetActivity(ctx context.Context, req dto.GetActivityRequest) ([]dto.GetActivityResponse, error) {
	valErr := a.validator.Validate(req)
	if valErr != nil {
		return []dto.GetActivityResponse{}, valErr
	}

	if req.Limit == 0 {
		req.Limit = 5
	}

	if req.Offset == 0 {
		req.Offset = 0
	}

	activities, err := a.activityRepo.GetActivity(ctx, enums.ActivityTypesReverse[req.ActivityType], req.DoneAtFrom, req.DoneAtTo, req.CaloriesBurnedMin, req.CaloriesBurnedMax, req.Limit, req.Offset)
	if err != nil {
		return []dto.GetActivityResponse{}, err
	}

	res := []dto.GetActivityResponse{}
	for _, activity := range activities {
		res = append(res, dto.GetActivityResponse{
			ActivityID:        activity.ID,
			AcitivityType:     enums.ActivityTypes[activity.ActivityType],
			DoneAt:            activity.DoneAt,
			DurationInMinutes: activity.DurationInMinutes,
			CaloriesBurned:    activity.CaloriesBurned,
			CreatedAt:         activity.CreatedAt,
		})
	}

	return res, nil
}

func (a *activityStruct) UpdateActivity(ctx context.Context, req dto.UpdateActivityRequest) (dto.UpdateActivityResponse, error) {
	log.Info(log.LogInfo{
		"req": req,
	}, "[UpdateActivity] Request")

	valErr := a.validator.Validate(req)
	if valErr != nil {
		return dto.UpdateActivityResponse{}, valErr
	}

	activity, err := a.activityRepo.GetActivityByID(ctx, req.ActivityID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dto.UpdateActivityResponse{}, domain.ErrNotFound
		}

		return dto.UpdateActivityResponse{}, err
	}

	if req.ActivityType == nil || req.DoneAt == nil || req.DurationInMinutes == nil {
		return dto.UpdateActivityResponse{}, fiber.NewError(fiber.StatusBadRequest, "Null is not accepted")
	}

	if *req.ActivityType != "" && *req.ActivityType != enums.ActivityTypes[activity.ActivityType] {
		activity.ActivityType = enums.ActivityTypesReverse[*req.ActivityType]
	}

	if !req.DoneAt.IsZero() && !(*req.DoneAt).Equal(activity.DoneAt) {
		activity.DoneAt = *req.DoneAt
	}

	if *req.DurationInMinutes != 0 && *req.DurationInMinutes != activity.DurationInMinutes {
		activity.DurationInMinutes = *req.DurationInMinutes
	}

	activity.CaloriesBurned = enums.Calories[enums.ActivityTypes[activity.ActivityType]] * activity.DurationInMinutes
	activity.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	err = a.activityRepo.UpdateActivity(ctx, activity)
	if err != nil {
		return dto.UpdateActivityResponse{}, err
	}

	doneAt := activity.DoneAt.UTC().Format("2006-01-02T15:04:05.000Z")
	res := dto.UpdateActivityResponse{
		ActivityID:        activity.ID,
		AcitivityType:     enums.ActivityTypes[activity.ActivityType],
		DoneAt:            doneAt,
		DurationInMinutes: activity.DurationInMinutes,
		CaloriesBurned:    activity.CaloriesBurned,
		CreatedAt:         activity.CreatedAt,
		UpdateAt:          activity.UpdatedAt.Time,
	}

	return res, nil
}

// DeleteActivity implements contracts.ActivityService.
func (a *activityStruct) DeleteActivity(ctx context.Context, req dto.DeleteActivityRequest) error {
	valErr := a.validator.Validate(req)
	if valErr != nil {
		return valErr
	}

	err := a.activityRepo.DeleteActivity(ctx, req.ActivityID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.ErrNotFound
		}

		return err
	}

	return nil
}
