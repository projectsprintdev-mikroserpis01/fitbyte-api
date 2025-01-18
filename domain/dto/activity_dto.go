package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateActivityRequest struct {
	ActivityType      string    `json:"activityType" validate:"required,oneof=Walking Yoga Stretching Cycling Swimming Dancing Hiking Running HIIT JumpRope"`
	DoneAt            time.Time `json:"doneAt" validate:"required"`
	DurationInMinutes int       `json:"durationInMinutes" validate:"required,gte=1"`
	UserID            int       `json:"userId" validate:"required"`
}

type CreateActivityResponse struct {
	ActivityID        uuid.UUID `json:"activityId"`
	AcitivityType     string    `json:"activityType"`
	DoneAt            time.Time `json:"doneAt"`
	DurationInMinutes int       `json:"durationInMinutes"`
	CaloriesBurned    int       `json:"caloriesBurned"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdateAt          time.Time `json:"updatedAt"`
}

type GetActivityRequest struct {
	Limit             int       `query:"limit" validate:"numeric"`
	Offset            int       `query:"offset" validate:"numeric"`
	ActivityType      string    `query:"activityType" validate:"omitempty,oneof=Walking Yoga Stretching Cycling Swimming Dancing Hiking Running HIIT JumpRope"`
	DoneAtFrom        time.Time `query:"doneAtFrom"`
	DoneAtTo          time.Time `query:"doneAtTo"`
	CaloriesBurnedMin int       `query:"caloriesBurnedMin" validate:"numeric"`
	CaloriesBurnedMax int       `query:"caloriesBurnedMax" validate:"numeric"`
}

type GetActivityResponse struct {
	ActivityID        uuid.UUID `json:"activityId"`
	AcitivityType     string    `json:"activityType"`
	DoneAt            time.Time `json:"doneAt"`
	DurationInMinutes int       `json:"durationInMinutes"`
	CaloriesBurned    int       `json:"caloriesBurned"`
	CreatedAt         time.Time `json:"createdAt"`
}

type UpdateActivityRequest struct {
	ActivityID        uuid.UUID  `param:"activityId" validate:"required,uuid"`
	ActivityType      *string    `json:"activityType" validate:"omitempty,oneof=Walking Yoga Stretching Cycling Swimming Dancing Hiking Running HIIT JumpRope"`
	DoneAt            *time.Time `json:"doneAt"`
	DurationInMinutes *int       `json:"durationInMinutes" validate:"gte=1"`
	UserID            int        `json:"userId" validate:"required"`
}

type UpdateActivityResponse struct {
	ActivityID        uuid.UUID `json:"activityId"`
	AcitivityType     string    `json:"activityType"`
	DoneAt            time.Time `json:"doneAt"`
	DurationInMinutes int       `json:"durationInMinutes"`
	CaloriesBurned    int       `json:"caloriesBurned"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdateAt          time.Time `json:"updatedAt"`
}

type DeleteActivityRequest struct {
	ActivityID uuid.UUID `param:"activityID" validate:"required,uuid"`
}
