package model

import "time"

type Category struct {
	ID        int64
	Name      string
	DomainID  int64
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type NewCategory struct {
	Name     string
	DomainID int64
}

type UpdatedCategory struct {
	Name     *string
	DomainID *int64
}
