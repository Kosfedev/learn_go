package model

import (
	"database/sql"
	"time"
)

type Question struct {
	ID              int64
	Text            string
	Type            int64
	ReferenceAnswer sql.NullString
	CreatedAt       time.Time
	UpdatedAt       sql.NullTime
}

type NewQuestion struct {
	Text            string
	Type            int64
	ReferenceAnswer sql.NullString
}

type UpdatedQuestion struct {
	Text            sql.NullString
	Type            sql.NullInt64
	ReferenceAnswer sql.NullString
}
