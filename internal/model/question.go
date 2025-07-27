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
	Id        int
	Text      string
	IsCorrect bool
}

type NewQuestionOption struct {
	Text      string
	IsCorrect bool
}

type Question struct {
	Id              int
	Text            string
	Type            QuestionType
	Options         *[]*QuestionOption
	ReferenceAnswer *string
	CreatedAt       time.Time
	UpdatedAt       *time.Time
}

type NewQuestion struct {
	Text            string
	Type            QuestionType
	Options         *[]*NewQuestionOption
	ReferenceAnswer *string
}

type UpdatedQuestion struct {
	Text            *string
	Type            *QuestionType
	Options         *[]*NewQuestionOption
	ReferenceAnswer *string
}
