package model

import (
	"database/sql"
	"time"
)

type Subcategory struct {
	ID         int64
	Name       string
	CategoryID int64
	CreatedAt  time.Time
	UpdatedAt  sql.NullTime
}

type NewSubcategory struct {
	Name       string
	CategoryID int64
}

type UpdatedSubcategory struct {
	Name       sql.NullString
	CategoryID sql.NullInt64
}
