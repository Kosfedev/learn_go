package model

type QuestionOption struct {
	ID         int64  `db:"id" json:"id"`
	QuestionID int64  `db:"question_id" json:"question_id"`
	Text       string `db:"text" json:"text"`
	IsCorrect  bool   `db:"is_correct" json:"is_correct"`
}

type NewQuestionOption struct {
	QuestionID int64  `db:"question_id" json:"question_id"`
	Text       string `db:"text" json:"text"`
	IsCorrect  bool   `db:"is_correct" json:"is_correct"`
}
