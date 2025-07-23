package model

import "time"

type QuestionSet struct {
	Id          int
	Name        string
	QuestionIds []int
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

type NewQuestionSet struct {
	Name        string
	QuestionIds []int
}

type UpdatedQuestionSet struct {
	Name        *string
	QuestionIds *[]int
}
