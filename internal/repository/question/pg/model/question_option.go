package model

type QuestionOption struct {
	ID         int32
	QuestionID int32
	Text       string
	IsCorrect  bool
}

type NewQuestionOption struct {
	QuestionID int32
	Text       string
	IsCorrect  bool
}
