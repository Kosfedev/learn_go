package pg

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/questionset/pg/converter"
)

func (r *repo) Create(ctx context.Context, newQuestionSet *model.NewQuestionSet) (int64, error) {
	newQuestionSetRepo := converter.NewQuestionSetToPGSQL(newQuestionSet)
	query := db.Query{
		Name:     "question_set_repository.create",
		QueryRaw: `INSERT INTO question_set (name) VALUES ($1) RETURNING id;`,
	}

	var questionSetIDRepo int32
	err := r.db.DB().ScanOne(ctx, &questionSetIDRepo, query, newQuestionSetRepo.Name)
	if err != nil {
		return 0, err
	}

	questionSetID := int64(questionSetIDRepo)

	// TODO: нужна транзакция
	err = r.addQuestions(ctx, questionSetID, newQuestionSet.QuestionIDs)
	if err != nil {
		return 0, err
	}

	return questionSetID, nil
}
