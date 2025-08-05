package converter

import (
	"github.com/Kosfedev/learn_go/internal/model"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/question/pg/model"
)

func QuestionOptionFromPGSQL(questionOptionRepo *modelRepo.QuestionOption) *model.QuestionOption {
	questionOptionServ := &model.QuestionOption{
		ID:        questionOptionRepo.ID,
		Text:      questionOptionRepo.Text,
		IsCorrect: questionOptionRepo.IsCorrect,
	}

	return questionOptionServ
}

func NewQuestionOptionToPGSQL(questionID int64, questionOptionServ *model.NewQuestionOption) *modelRepo.NewQuestionOption {
	newQuestionOptionRepo := &modelRepo.NewQuestionOption{
		QuestionID: questionID,
		Text:       questionOptionServ.Text,
		IsCorrect:  questionOptionServ.IsCorrect,
	}

	return newQuestionOptionRepo
}
