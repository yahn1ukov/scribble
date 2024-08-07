package model

import "time"

type User struct {
	ID        string    `db:"id"`
	Email     string    `db:"email"`
	FirstName *string   `db:"first_name"`
	LastName  *string   `db:"last_name"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
