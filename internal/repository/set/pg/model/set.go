package model

import (
	"database/sql"
	"time"
)

type Set struct {
	ID        int64        `db:"id" json:"id"`
	Name      string       `db:"name" json:"name"`
	CreatedAt time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at" json:"updated_at"`
}

type NewSet struct {
	Name string `db:"name" json:"name"`
}

type UpdatedSet struct {
	Name sql.NullString `db:"name" json:"name"`
}

type SetListWithTotal struct {
	Sets  []*Set `json:"sets"`
	Total int64  `db:"total" json:"total"`
}
