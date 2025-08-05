package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/client/db"
)

func (r *repo) RemoveSubcategoriesFromQuestion(ctx context.Context, questionID int64, subcategoryIDs []int64) error {
	if len(subcategoryIDs) == 0 {
		return nil
	}

	values := make([]interface{}, len(subcategoryIDs)+1)
	values[0] = questionID
	placeholders := make([]string, len(subcategoryIDs))
	for i, id := range subcategoryIDs {
		values[i+1] = id
		placeholders[i] = fmt.Sprintf("$%d", i+2)
	}

	query := db.Query{
		Name: "question_subcategory_repository.remove_subcategories_from_question",
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
