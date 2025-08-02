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

func NewQuestionOptionToPGSQL(questionId int, questionOptionServ *model.NewQuestionOption) *modelRepo.NewQuestionOption {
	newQuestionOptionRepo := &modelRepo.NewQuestionOption{
		QuestionId: int32(questionId),
		Text:       questionOptionServ.Text,
		IsCorrect:  questionOptionServ.IsCorrect,
	}

	return newQuestionOptionRepo
}
