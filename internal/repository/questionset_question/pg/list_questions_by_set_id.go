package pg

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/repository/question_subcategory/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/question_subcategory/pg/model"
)

func (r *repo) ListQuestionsBySetID(ctx context.Context, questionSetID int64) ([]int64, error) {
	query := db.Query{
		Name:     "question_subcategory_repository.list_questions_by_set_id",
		QueryRaw: `SELECT * FROM question_set_question WHERE question_set_id = $1`,
	}

	questionSets := make([]*modelRepo.QuestionSubcategory, 0)
	err := r.db.DB().ScanAll(ctx, &questionSets, query, questionSetID)
	if err != nil {
		return nil, err
	}

	questionIDs := converter.QuestionSubcategoriesIDsFromPGSQL(questionSets)

	return questionIDs, nil
}
