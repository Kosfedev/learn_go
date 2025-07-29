package model

type QuestionOption struct {
	Id         int32
	QuestionId int32
	Text       string
	IsCorrect  bool
}

type NewQuestionOption struct {
	Text      string
	IsCorrect bool
}
