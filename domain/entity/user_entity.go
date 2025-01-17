package entity

import "time"

type User struct {
	ID         int       `db:"id"`
	Email      string    `db:"email"`
	Password   string    `db:"password"`
	Name       string    `db:"name"`
	Preference string    `db:"preference"`
	WeightUnit string    `db:"weight_unit"`
	HeightUnit string    `db:"height_unit"`
	Weight     int       `db:"weight"`
	Height     int       `db:"height"`
	ImageURI   string    `db:"image_uri"`
	CreatedAt  time.Time `db:"created_at"`
}
