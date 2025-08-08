package model

import (
	"database/sql"
	"time"
)

type Subcategory struct {
	ID         int64        `db:"id" json:"id"`
	Name       string       `db:"name" json:"name"`
	CategoryID int64        `db:"category_id" json:"category_id"`
	CreatedAt  time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt  sql.NullTime `db:"updated_at" json:"updated_at"`
}

type NewSubcategory struct {
	Name       string `db:"name" json:"name"`
	CategoryID int64  `db:"category_id" json:"category_id"`
}

type UpdatedSubcategory struct {
	Name       sql.NullString `db:"name" json:"name"`
	CategoryID sql.NullInt64  `db:"category_id" json:"category_id"`
}
