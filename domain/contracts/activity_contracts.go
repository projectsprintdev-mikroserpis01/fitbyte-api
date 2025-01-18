package contracts

import (
	"context"
	"time"

	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/dto"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/entity"
)

type ActivityRepository interface {
	CreateActivity(ctx context.Context, activity entity.Activity) error
	GetActivity(ctx context.Context, activityType int16, doneAtFrom, doneAtTo time.Time, caloriesBurnedMin, caloriesBurnedMax, limit, offset int) ([]entity.Activity, error)
	UpdateActivity(ctx context.Context, activity entity.Activity) error
	DeleteActivity(ctx context.Context, activityID string) error
}

type ActivityService interface {
	CreateActivity(ctx context.Context, req dto.CreateActivityRequest) (dto.CreateActivityResponse, error)
	GetActivity(ctx context.Context, req dto.GetActivityRequest) ([]dto.GetActivityResponse, error)
	UpdateActivity(ctx context.Context, req dto.UpdateActivityRequest) (dto.UpdateActivityResponse, error)
	DeleteActivity(ctx context.Context, req dto.DeleteActivityRequest) error
}
