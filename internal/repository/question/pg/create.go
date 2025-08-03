package pg

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/question/pg/converter"
)

func (r *repo) Create(ctx context.Context, newQuestion *model.NewQuestion) (int, error) {
	newQuestionRepo := converter.NewQuestionToPGSQL(newQuestion)
	query := db.Query{
		Name: "question_repository.create",
		QueryRaw: `
			INSERT INTO question (text, type, reference_answer) 
			VALUES ($1, $2, $3)
			RETURNING id;`,
	}

	var questionId int
	err := r.db.DB().ScanOne(ctx, &questionId, query, newQuestionRepo.Text, newQuestionRepo.Type, newQuestionRepo.ReferenceAnswer)
	if err != nil {
		return 0, err
	}

	// TODO: нужна транзакция
	err = r.AddOptions(ctx, questionId, newQuestion.Options)

	return questionId, nil
}
