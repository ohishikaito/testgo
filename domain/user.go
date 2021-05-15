package domain

import "time"

type User struct {
	Id        string
	LastName  string
	FirstName string
	CreatedAt time.Time
	UpdatedAt time.Time
}
