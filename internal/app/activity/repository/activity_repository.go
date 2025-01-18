package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/contracts"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/domain/entity"
	"github.com/projectsprintdev-mikroserpis01/fitbyte-api/pkg/log"
)

type activityRepo struct {
	db *sqlx.DB
}

func NewActivityRepository(db *sqlx.DB) contracts.ActivityRepository {
	return &activityRepo{
		db: db,
	}
}

func (a *activityRepo) CreateActivity(ctx context.Context, activity entity.Activity) error {
	_, err := a.db.ExecContext(ctx, "INSERT INTO activities (id, activity_type, done_at, duration_in_minutes, user_id, calories_burned) VALUES ($1, $2, $3, $4, $5, $6)", activity.ID, activity.ActivityType, activity.DoneAt, activity.DurationInMinutes, activity.UserID, activity.CaloriesBurned)
	if err != nil {
		return err
	}

	return nil
}

func (a *activityRepo) GetActivity(ctx context.Context, activityType int16, doneAtFrom time.Time, doneAtTo time.Time, caloriesBurnedMin int, caloriesBurnedMax int, limit int, offset int) ([]entity.Activity, error) {
	activities := []entity.Activity{}

	query := "SELECT * FROM activities WHERE 1=1"
	params := []interface{}{}

	// Apply filters
	if activityType != 0 {
		query += fmt.Sprintf(" AND activity_type = $%d", len(params)+1)
		params = append(params, activityType)
	}

	if !doneAtFrom.IsZero() {
		query += fmt.Sprintf(" AND done_at >= $%d", len(params)+1)
		params = append(params, doneAtFrom)
	}

	if !doneAtTo.IsZero() {
		query += fmt.Sprintf(" AND done_at <= $%d", len(params)+1)
		params = append(params, doneAtTo)
	}

	if caloriesBurnedMin > 0 {
		query += fmt.Sprintf(" AND calories_burned >= $%d", len(params)+1)
		params = append(params, caloriesBurnedMin)
	}

	if caloriesBurnedMax > 0 {
		query += fmt.Sprintf(" AND calories_burned <= $%d", len(params)+1)
		params = append(params, caloriesBurnedMax)
	}

	query += fmt.Sprintf(" ORDER BY done_at DESC LIMIT $%d OFFSET $%d", len(params)+1, len(params)+2)
	params = append(params, limit, offset)

	log.Info(log.LogInfo{
		"query":  query,
		"params": params,
	}, "[GetActivity] Query")

	err := a.db.SelectContext(ctx, &activities, query, params...)
	if err != nil {
		return nil, err
	}

	return activities, nil
}

func (a *activityRepo) GetActivityByID(ctx context.Context, activityID uuid.UUID) (entity.Activity, error) {
	activity := entity.Activity{}

	err := a.db.GetContext(ctx, &activity, "SELECT * FROM activities WHERE id = $1", activityID)
	if err != nil {
		return entity.Activity{}, err
	}

	return activity, nil
}


func (a *activityRepo) UpdateActivity(ctx context.Context, activity entity.Activity) error {
	res, err := a.db.ExecContext(ctx, "UPDATE activities SET activity_type = $1, done_at = $2, duration_in_minutes = $3, calories_burned = $4, updated_at = $5 WHERE id = $6", activity.ActivityType, activity.DoneAt, activity.DurationInMinutes, activity.CaloriesBurned, activity.UpdatedAt, activity.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domain.ErrEntityNotFound
	} else if rowsAffected > 1 {
		return domain.ErrMultipleEntities
	}

	return nil
}

func (a *activityRepo) DeleteActivity(ctx context.Context, activityID uuid.UUID) error {
	res, err := a.db.ExecContext(ctx, "DELETE FROM activities WHERE id = $1", activityID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domain.ErrEntityNotFound
	} else if rowsAffected > 1 {
		return domain.ErrMultipleEntities
	}

	return nil
}
