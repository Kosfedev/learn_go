package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/client/db"
)

func (r *repo) RemoveQuestionFromSets(ctx context.Context, questionID int64, questionSetIDs []int64) error {
	if len(questionSetIDs) == 0 {
		return nil
	}

	values := make([]interface{}, len(questionSetIDs)+1)
	values[0] = questionID
	placeholders := make([]string, len(questionSetIDs))
	for i, id := range questionSetIDs {
		values[i+1] = id
		placeholders[i] = fmt.Sprintf("$%d", i+2)
	}

	query := db.Query{
		Name: "question_subcategory_repository.remove_question_from_sets",
		QueryRaw: fmt.Sprintf(
			"DELETE FROM question_set_question WHERE question_id = $1 AND question_set_id IN (%s)",
			strings.Join(placeholders, ", "),
		),
	}

	_, err := r.db.DB().ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}
