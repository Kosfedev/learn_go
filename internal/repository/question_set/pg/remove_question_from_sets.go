package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/client/db"
)

func (r *repo) RemoveQuestionFromSets(ctx context.Context, questionID int64, setIDs []int64) error {
	if len(setIDs) == 0 {
		return nil
	}

	values := make([]interface{}, len(setIDs)+1)
	values[0] = questionID
	placeholders := make([]string, len(setIDs))
	for i, id := range setIDs {
		values[i+1] = id
		placeholders[i] = fmt.Sprintf("$%d", i+2)
	}

	query := db.Query{
		Name: "question_set_repository.remove_question_from_sets",
		QueryRaw: fmt.Sprintf(
			"DELETE FROM question_set WHERE question_id = $1 AND set_id IN (%s)",
			strings.Join(placeholders, ", "),
		),
	}

	_, err := r.db.DB().ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}
