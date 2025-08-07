package pg

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/repository/question_set/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/question_set/pg/model"
)

func (r *repo) ListSetsByQuestionID(ctx context.Context, questionID int64) ([]int64, error) {
	query := db.Query{
		Name:     "question_set_repository.list_sets_by_question_id",
		QueryRaw: `SELECT * FROM question_set WHERE question_id = $1`,
	}

	questionSets := make([]*modelRepo.QuestionSet, 0)
	err := r.db.DB().ScanAll(ctx, &questionSets, query, questionID)
	if err != nil {
		return nil, err
	}

	setIDs := converter.QuestionSetsIDsFromPGSQL(questionSets)

	return setIDs, nil
}
