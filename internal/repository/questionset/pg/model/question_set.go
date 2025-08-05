package model

import (
	"database/sql"
	"time"
)

type QuestionSet struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type NewQuestionSet struct {
	Name string
}

type UpdatedQuestionSet struct {
	Name sql.NullString
}
