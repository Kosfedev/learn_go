package model

import (
	"database/sql"
	"time"
)

type Subcategory struct {
	Id         int32
	Name       string
	CategoryId int32
	CreatedAt  time.Time
	UpdatedAt  sql.NullTime
}

type NewSubcategory struct {
	Name       string
	CategoryId int32
}

type UpdatedSubcategory struct {
	Name       sql.NullString
	CategoryId sql.NullInt32
}
