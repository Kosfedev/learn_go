package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/client/db"
)

func (r *repo) AddSubcategories(ctx context.Context, questionId int, subcategoryIds []int) error {
	if len(subcategoryIds) == 0 {
		return nil
	}

	var values []interface{}
	nArgs := 2
	questionIdRepo := int32(questionId)
	queryRaw := `INSERT INTO question_subcategory (question_id, subcategory_id) VALUES`

	for i, subcategoryId := range subcategoryIds {
		queryRaw += fmt.Sprintf(" ($%d, $%d),", i*nArgs+1, i*nArgs+2)
		values = append(values, questionIdRepo, int32(subcategoryId))
	}

	queryRaw = strings.TrimSuffix(queryRaw, ",")
	query := db.Query{
		Name:     "question_repository.add_subcategories",
		QueryRaw: queryRaw,
	}

	_, err := r.db.DB().ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}
