package model

import "time"

const (
	SingleAnswer = iota
	MultiAnswer
	OpenAnswer
)

type QuestionType uint8

func (qt QuestionType) isValid() bool {
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
	Text    string
	Correct bool
}

type Question struct {
	Id              int
	Text            string
	Type            QuestionType
	Options         *[]QuestionOption
	ReferenceAnswer *string
	CreatedAt       time.Time
	UpdatedAt       *time.Time
}

type NewQuestion struct {
	Text            string
	Type            QuestionType
	Options         *[]QuestionOption
	ReferenceAnswer *string
}

type UpdatedQuestion struct {
	Text            *string
	Type            *QuestionType
	Options         *[]QuestionOption
	ReferenceAnswer *string
}

/*
Rest of base types
type QuestionsDomain struct {
	Id   int
	Name string
}

type QuestionsCategory struct {
	Id       int
	Name     string
	DomainId int
}

type QuestionsSubCategory struct {
	Id         int
	Name       string
	CategoryId int
}

type QuestionsSet struct {
	Id   int
	Name string
}

type QuestionsSetItems struct {
	QuestionId int
	SetId      int
}

type QuestionSubcategory struct {
	QuestionId    int
	SubcategoryId int
}*/
