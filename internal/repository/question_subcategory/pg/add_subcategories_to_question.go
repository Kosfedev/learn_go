package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/client/db"
)

func (r *repo) AddSubcategoriesToQuestion(ctx context.Context, questionID int, subcategoryIDs []int) error {
	if len(subcategoryIDs) == 0 {
		return nil
	}

	var values []interface{}
	nArgs := 2
	questionIDRepo := int32(questionID)
	queryRaw := `INSERT INTO question_subcategory (question_id, subcategory_id) VALUES`

	for i, subcategoryID := range subcategoryIDs {
		queryRaw += fmt.Sprintf(" ($%d, $%d),", i*nArgs+1, i*nArgs+2)
		values = append(values, questionIDRepo, int32(subcategoryID))
	}

	queryRaw = strings.TrimSuffix(queryRaw, ",")
	query := db.Query{
		Name:     "question_subcategory_repository.add_subcategories_to_question",
		QueryRaw: queryRaw,
	}

	_, err := r.db.DB().ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}
