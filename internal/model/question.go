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

type QuestionOption struct {
	ID        int64
	Text      string
	IsCorrect bool
}

type NewQuestionOption struct {
	Text      string
	IsCorrect bool
}

type Question struct {
	ID              int64
	Text            string
	Type            QuestionType
	Options         []*QuestionOption
	Subcategories   []*Subcategory
	Sets            []*Set
	ReferenceAnswer *string
	CreatedAt       time.Time
	UpdatedAt       *time.Time
}

type NewQuestion struct {
	Text            string
	Type            QuestionType
	ReferenceAnswer *string
	Options         []*NewQuestionOption
}

type UpdatedQuestion struct {
	Text            *string
	Type            *QuestionType
	ReferenceAnswer *string
}
