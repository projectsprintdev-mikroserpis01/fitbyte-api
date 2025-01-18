package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Activity struct {
	ID                uuid.UUID    `db:"id"`
	ActivityType      int16        `db:"activity_type"`
	DoneAt            time.Time    `db:"done_at"`
	DurationInMinutes int          `db:"duration_in_minutes"`
	UserID            uuid.UUID    `db:"user_id"`
	CaloriesBurned    int          `db:"calories_burned"`
	CreatedAt         time.Time    `db:"created_at"`
	UpdatedAt         sql.NullTime `db:"updated_at"`
}
