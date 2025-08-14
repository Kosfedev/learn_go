package converter

import (
	"database/sql"

	"github.com/Kosfedev/learn_go/internal/model"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/question/pg/model"
)

func QuestionsFromPGSQL(questionRepo []*modelRepo.Question) []*model.Question {
	questions := make([]*model.Question, len(questionRepo))
	for i, question := range questionRepo {
		questions[i] = QuestionFromPGSQL(question)
	}

	return questions
}

func QuestionFromPGSQL(questionRepo *modelRepo.Question) *model.Question {
	question := &model.Question{
		ID:        questionRepo.ID,
		Text:      questionRepo.Text,
		Type:      model.QuestionType(questionRepo.Type),
		CreatedAt: questionRepo.CreatedAt,
	}

	if questionRepo.ReferenceAnswer.Valid {
		question.ReferenceAnswer = &questionRepo.ReferenceAnswer.String
	}

	if questionRepo.UpdatedAt.Valid {
		question.UpdatedAt = &questionRepo.UpdatedAt.Time
	}

	return question
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

func UpdatedQuestionToPGSQL(question *model.UpdatedQuestion) *modelRepo.UpdatedQuestion {
	questionRepo := &modelRepo.UpdatedQuestion{}

	if question.Text != nil {
		questionRepo.Text = sql.NullString{
			String: *question.Text,
			Valid:  true,
		}
	}

	if question.Type != nil {
		questionRepo.Type = sql.NullInt64{
			Int64: int64(*question.Type),
			Valid: true,
		}
	}

	if question.ReferenceAnswer != nil {
		questionRepo.ReferenceAnswer = sql.NullString{
			String: *question.ReferenceAnswer,
			Valid:  true,
		}
	}

	return questionRepo
}
