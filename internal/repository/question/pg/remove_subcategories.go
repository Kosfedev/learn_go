package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/client/db"
)

func (r *repo) RemoveSubcategories(ctx context.Context, questionId int, subcategoryIds []int) error {
	if len(subcategoryIds) == 0 {
		return nil
	}

	values := make([]interface{}, len(subcategoryIds)+1)
	values[0] = questionId
	placeholders := make([]string, len(subcategoryIds))
	for i, id := range subcategoryIds {
		values[i+1] = int32(id)
		placeholders[i] = fmt.Sprintf("$%d", i+2)
	}

	query := db.Query{
		Name: "question_repository.remove_subcategories",
		QueryRaw: fmt.Sprintf(
			"DELETE FROM question_subcategory WHERE question_id = $1 AND subcategory_id IN (%s)",
			strings.Join(placeholders, ", "),
		),
	}
	fmt.Printf("%s\n", query.QueryRaw)
	fmt.Printf("%+v\n", values)
	_, err := r.db.DB().ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}
