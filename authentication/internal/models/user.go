package models

import "time"

type User struct {
	ID             uint
	Name           string
	Surname        string
	Email          string
	Country        string
	HashedPassword string
	CreatedAt      time.Time
}
