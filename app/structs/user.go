package structs

import (
	"time"
)

type User struct {
	ID           int
	Name         string
	Occupation   string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	RoleID       int
	FileID       *int
	Role         Role
	File         File
}
