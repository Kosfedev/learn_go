package pg

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/repository/question_set/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/question_set/pg/model"
)

func (r *repo) ListQuestionsBySetID(ctx context.Context, setID int64) ([]int64, error) {
	query := db.Query{
		Name:     "question_set_repository.list_questions_by_set_id",
		QueryRaw: `SELECT * FROM question_set WHERE set_id = $1`,
	}

	questionSets := make([]*modelRepo.QuestionSet, 0)
	err := r.db.DB().ScanAll(ctx, &questionSets, query, setID)
	if err != nil {
		return nil, err
	}

	questionIDs := converter.QuestionIDsFromPGSQL(questionSets)

	return questionIDs, nil
}
