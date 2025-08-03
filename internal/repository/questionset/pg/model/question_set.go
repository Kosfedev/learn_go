package model

import (
	"database/sql"
	"time"
)

type QuestionSet struct {
	Id        int32
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
