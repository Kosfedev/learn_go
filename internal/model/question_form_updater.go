package model

type NewQuestionWithOptions struct {
	Question *NewQuestion
	Options  []*NewQuestionOption
}
