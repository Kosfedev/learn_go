package model

import (
	"database/sql"
	"time"
)

type Question struct {
	Id              int32
	Text            string
	Type            int32
	ReferenceAnswer sql.NullString
	CreatedAt       time.Time
	UpdatedAt       sql.NullTime
}

type NewQuestion struct {
	Text            string
	Type            int32
	ReferenceAnswer sql.NullString
}

type UpdatedQuestion struct {
	Text            sql.NullString
	Type            sql.NullInt32
	ReferenceAnswer sql.NullString
}
