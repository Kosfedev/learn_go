package model

import "time"

type QuestionSet struct {
	ID          int
	Name        string
	QuestionIDs []int
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

type NewQuestionSet struct {
	Name        string
	QuestionIDs []int
}

type UpdatedQuestionSet struct {
	Name *string
}
