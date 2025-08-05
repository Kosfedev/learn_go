package model

import "time"

type Domain struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type NewDomain struct {
	Name string
}

type UpdatedDomain struct {
	Name *string
}
