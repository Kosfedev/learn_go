package pg

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/repository/question/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/question/pg/model"
)

func (r *repo) ListSubcategoriesByQuestionId(ctx context.Context, questionId int) ([]int, error) {
	idRepo := int32(questionId)
	query := db.Query{
		Name:     "subcategory_repository.list_by_ids",
		QueryRaw: `SELECT * FROM question_subcategory WHERE question_id = $1`,
	}

	questionSubcategories := make([]*modelRepo.QuestionSubcategory, 0)
	err := r.db.DB().ScanAll(ctx, &questionSubcategories, query, idRepo)
	if err != nil {
		return nil, err
	}

	subcategories := converter.QuestionSubcategoriesIdsFromPGSQL(questionSubcategories)

	return subcategories, nil
}
