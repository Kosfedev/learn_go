package model

import "time"

type Subcategory struct {
	Id         int
	Name       string
	CategoryId int
	CreatedAt  time.Time
	UpdatedAt  *time.Time
}

type NewSubcategory struct {
	Name       string
	CategoryId int
}

type UpdatedSubcategory struct {
	Name       *string
	CategoryId *int
}
