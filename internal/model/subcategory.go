package model

import "time"

type Subcategory struct {
	ID         int64
	Name       string
	CategoryID int64
	CreatedAt  time.Time
	UpdatedAt  *time.Time
}

type NewSubcategory struct {
	Name       string
	CategoryID int64
}

type UpdatedSubcategory struct {
	Name       *string
	CategoryID *int64
}
