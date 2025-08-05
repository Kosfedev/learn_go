package model

import (
	"database/sql"
	"time"
)

type Subcategory struct {
	ID         int32
	Name       string
	CategoryID int32
	CreatedAt  time.Time
	UpdatedAt  sql.NullTime
}

type NewSubcategory struct {
	Name       string
	CategoryID int32
}

type UpdatedSubcategory struct {
	Name       sql.NullString
	CategoryID sql.NullInt32
}
