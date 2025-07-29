package converter

import (
	"github.com/Kosfedev/learn_go/internal/model"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/question/pg/model"
)

func QuestionOptionFromPGSQL(questionOptionRepo *modelRepo.QuestionOption) *model.QuestionOption {
	questionOptionServ := &model.QuestionOption{
		Id:        int(questionOptionRepo.Id),
		Text:      questionOptionRepo.Text,
		IsCorrect: questionOptionRepo.IsCorrect,
	}

	return questionOptionServ
}

func NewQuestionOptionToPGSQL(questionOptionServ *model.NewQuestionOption) *modelRepo.NewQuestionOption {
	questionOptionRepo := &modelRepo.NewQuestionOption{
		Text:      questionOptionServ.Text,
		IsCorrect: questionOptionServ.IsCorrect,
	}

	return questionOptionRepo
}
