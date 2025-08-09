package model

type QuestionOption struct {
	ID        int64
	Text      string
	IsCorrect bool
}

type NewQuestionOption struct {
	Text      string
	IsCorrect bool
}
