package model

import "time"

const (
	SingleAnswer = iota + 1
	MultiAnswer
	OpenAnswer
)

type QuestionType uint8

func (qt QuestionType) IsValid() bool {
	if qt == SingleAnswer || qt == MultiAnswer || qt == OpenAnswer {
		return true
	}

	return false
}

type Question struct {
	ID              int64
	Text            string
	Type            QuestionType
	ReferenceAnswer *string
	CreatedAt       time.Time
	UpdatedAt       *time.Time
}

type NewQuestion struct {
	Text            string
	Type            QuestionType
	ReferenceAnswer *string
}

type UpdatedQuestion struct {
	Text            *string
	Type            *QuestionType
	ReferenceAnswer *string
}
