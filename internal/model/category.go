package model

import "time"

type Category struct {
	ID        int
	Name      string
	DomainID  int
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type NewCategory struct {
	Name     string
	DomainID int
}

type UpdatedCategory struct {
	Name     *string
	DomainID *int
}
