package model

import (
	"database/sql"
	"time"
)

type Question struct {
	ID              int64          `db:"id" json:"id"`
	Text            string         `db:"text" json:"text"`
	Type            int64          `db:"type" json:"type"`
	ReferenceAnswer sql.NullString `db:"reference_answer" json:"reference_answer"`
	CreatedAt       time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt       sql.NullTime   `db:"updated_at" json:"updated_at"`
}

type NewQuestion struct {
	Text            string         `db:"text" json:"text"`
	Type            int64          `db:"type" json:"type"`
	ReferenceAnswer sql.NullString `db:"reference_answer" json:"reference_answer"`
}

type UpdatedQuestion struct {
	Text            sql.NullString `db:"text" json:"text"`
	Type            sql.NullInt64  `db:"type" json:"type"`
	ReferenceAnswer sql.NullString `db:"reference_answer" json:"reference_answer"`
}
