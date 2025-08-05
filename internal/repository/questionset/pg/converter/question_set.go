package converter

import (
	"database/sql"

	"github.com/Kosfedev/learn_go/internal/model"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/questionset/pg/model"
)

func QuestionSetFromPGSQL(questionSetRepo *modelRepo.QuestionSet) *model.QuestionSet {
	questionSetServ := &model.QuestionSet{
		ID:        questionSetRepo.ID,
		Name:      questionSetRepo.Name,
		CreatedAt: questionSetRepo.CreatedAt,
	}

	if questionSetRepo.UpdatedAt.Valid {
		questionSetServ.UpdatedAt = &questionSetRepo.UpdatedAt.Time
	}

	return questionSetServ
}

func NewQuestionSetToPGSQL(questionSetServ *model.NewQuestionSet) *modelRepo.NewQuestionSet {
	questionSetRepo := &modelRepo.NewQuestionSet{
		Name: questionSetServ.Name,
	}

	return questionSetRepo
}

func UpdatedQuestionSetToPGSQL(questionSetServ *model.UpdatedQuestionSet) *modelRepo.UpdatedQuestionSet {
	questionSetRepo := &modelRepo.UpdatedQuestionSet{}

	if questionSetServ.Name != nil {
		questionSetRepo.Name = sql.NullString{
			String: *questionSetServ.Name,
			Valid:  true,
		}
	}

	return questionSetRepo
}
