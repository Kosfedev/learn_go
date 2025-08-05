package converter

import (
	"database/sql"

	"github.com/Kosfedev/learn_go/internal/model"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/question/pg/model"
)

func QuestionFromPGSQL(questionRepo *modelRepo.Question) *model.Question {
	questionServ := &model.Question{
		ID:        questionRepo.ID,
		Text:      questionRepo.Text,
		Type:      model.QuestionType(questionRepo.Type),
		CreatedAt: questionRepo.CreatedAt,
	}

	if questionRepo.ReferenceAnswer.Valid {
		questionServ.ReferenceAnswer = &questionRepo.ReferenceAnswer.String
	}

	if questionRepo.UpdatedAt.Valid {
		questionServ.UpdatedAt = &questionRepo.UpdatedAt.Time
	}

	return questionServ
}

func NewQuestionToPGSQL(questionServ *model.NewQuestion) *modelRepo.NewQuestion {
	questionRepo := &modelRepo.NewQuestion{
		Text: questionServ.Text,
		Type: int64(questionServ.Type),
	}

	if questionServ.ReferenceAnswer != nil {
		questionRepo.ReferenceAnswer = sql.NullString{
			String: *questionServ.ReferenceAnswer,
			Valid:  true,
		}
	}

	return questionRepo
}

func UpdatedQuestionToPGSQL(questionServ *model.UpdatedQuestion) *modelRepo.UpdatedQuestion {
	questionRepo := &modelRepo.UpdatedQuestion{}

	if questionServ.Text != nil {
		questionRepo.Text = sql.NullString{
			String: *questionServ.Text,
			Valid:  true,
		}
	}

	if questionServ.Type != nil {
		questionRepo.Type = sql.NullInt64{
			Int64: int64(*questionServ.Type),
			Valid: true,
		}
	}

	if questionServ.ReferenceAnswer != nil {
		questionRepo.ReferenceAnswer = sql.NullString{
			String: *questionServ.ReferenceAnswer,
			Valid:  true,
		}
	}

	return questionRepo
}
