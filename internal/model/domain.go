package model

import "time"

type Domain struct {
	Id        int
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
