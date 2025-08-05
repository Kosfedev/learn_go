package pg

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/repository/question_subcategory/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/question_subcategory/pg/model"
)

func (r *repo) ListSubcategoriesByQuestionID(ctx context.Context, questionID int64) ([]int64, error) {
	idRepo := int32(questionID)
	query := db.Query{
		Name:     "question_subcategory_repository.list_subcategories_by_question_id",
		QueryRaw: `SELECT * FROM question_subcategory WHERE question_id = $1`,
	}

	questionSubcategories := make([]*modelRepo.QuestionSubcategory, 0)
	err := r.db.DB().ScanAll(ctx, &questionSubcategories, query, idRepo)
	if err != nil {
		return nil, err
	}

	subcategories := converter.QuestionSubcategoriesIDsFromPGSQL(questionSubcategories)

	return subcategories, nil
}
