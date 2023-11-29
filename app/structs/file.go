package structs

import "time"

type File struct {
	ID        int
	Name      string
	Location  string
	CreatedAt time.Time
	UpdatedAt time.Time
}