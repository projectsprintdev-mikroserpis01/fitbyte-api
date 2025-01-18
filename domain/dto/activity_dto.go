package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateActivityRequest struct {
	ActivityType      string    `json:"activityType" validate:"required,oneof=Walking Yoga Stretching Cylcing Swimming Dancing Hiking Runing HIIT JumpRope"`
	DoneAt            time.Time `json:"doneAt" validate:"required,datetime"`
	DurationInMinutes int       `json:"durationInMinutes" validate:"required,gte=1"`
}

type CreateActivityResponse struct {
	ActivityID        uuid.UUID `json:"activityId"`
	AcitivityType     string    `json:"activityType" validate:"required,oneof=Walking Yoga Stretching Cylcing Swimming Dancing Hiking Runing HIIT JumpRope"`
	DoneAt            time.Time `json:"doneAt"`
	DurationInMinutes int       `json:"durationInMinutes"`
	CaloriesBurned    int       `json:"caloriesBurned"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdateAt          time.Time `json:"updatedAt"`
}

type GetActivityRequest struct {
	Limit             int       `query:"limit" validate:"numeric"`
	Offset            int       `query:"offset" validate:"numeric"`
	DoneAtFrom        time.Time `query:"doneAtFrom" validate:"datetime"`
	DoneAtTo          time.Time `query:"doneAtTo" validate:"datetime"`
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
	ActivityType      string    `json:"activityType" validate:"oneof=Walking Yoga Stretching Cylcing Swimming Dancing Hiking Runing HIIT JumpRope"`
	DoneAt            time.Time `json:"doneAt" validate:"datetime"`
	DurationInMinutes int       `json:"durationInMinutes" validate:"gte=1"`
}

type UpdateActivityResponse struct {
	ActivityID        uuid.UUID `param:"activityId" validate:"required,uuid"`
	AcitivityType     string    `json:"activityType"`
	DoneAt            time.Time `json:"doneAt"`
	DurationInMinutes int       `json:"durationInMinutes"`
	CaloriesBurned    int       `json:"caloriesBurned"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdateAt          time.Time `json:"updatedAt"`
}

type DeleteActivityRequest struct {
	ActivityID uuid.UUID `param:"activityId" validate:"required,uuid"`
}
