package pg

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/repository/question_subcategory/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/question_subcategory/pg/model"
)

func (r *repo) ListSetsByQuestionID(ctx context.Context, questionID int64) ([]int64, error) {
	query := db.Query{
		Name:     "question_subcategory_repository.list_sets_by_question_id",
		QueryRaw: `SELECT * FROM question_set_question WHERE question_id = $1`,
	}

	questionSets := make([]*modelRepo.QuestionSubcategory, 0)
	err := r.db.DB().ScanAll(ctx, &questionSets, query, questionID)
	if err != nil {
		return nil, err
	}

	setIDs := converter.QuestionSubcategoriesIDsFromPGSQL(questionSets)

	return setIDs, nil
}
