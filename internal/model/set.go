package model

import "time"

type Set struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type NewSet struct {
	Name string
}

type UpdatedSet struct {
	Name *string
}
