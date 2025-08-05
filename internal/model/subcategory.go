package model

import "time"

type Subcategory struct {
	ID         int
	Name       string
	CategoryID int
	CreatedAt  time.Time
	UpdatedAt  *time.Time
}

type NewSubcategory struct {
	Name       string
	CategoryID int
}

type UpdatedSubcategory struct {
	Name       *string
	CategoryID *int
}
