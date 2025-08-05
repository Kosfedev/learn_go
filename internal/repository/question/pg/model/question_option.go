package model

type QuestionOption struct {
	ID         int64
	QuestionID int64
	Text       string
	IsCorrect  bool
}

type NewQuestionOption struct {
	QuestionID int64
	Text       string
	IsCorrect  bool
}
