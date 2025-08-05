package converter

import (
	"github.com/Kosfedev/learn_go/internal/model"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/question/pg/model"
)

func QuestionOptionFromPGSQL(questionOptionRepo *modelRepo.QuestionOption) *model.QuestionOption {
	questionOptionServ := &model.QuestionOption{
		ID:        int(questionOptionRepo.ID),
		Text:      questionOptionRepo.Text,
		IsCorrect: questionOptionRepo.IsCorrect,
	}

	return questionOptionServ
}

func NewQuestionOptionToPGSQL(questionID int, questionOptionServ *model.NewQuestionOption) *modelRepo.NewQuestionOption {
	newQuestionOptionRepo := &modelRepo.NewQuestionOption{
		QuestionID: int32(questionID),
		Text:       questionOptionServ.Text,
		IsCorrect:  questionOptionServ.IsCorrect,
	}

	return newQuestionOptionRepo
}
