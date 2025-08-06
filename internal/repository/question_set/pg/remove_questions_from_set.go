package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/client/db"
)

func (r *repo) RemoveQuestionsFromSet(ctx context.Context, setID int64, questionIDs []int64) error {
	if len(questionIDs) == 0 {
		return nil
	}

	values := make([]interface{}, len(questionIDs)+1)
	values[0] = setID
	placeholders := make([]string, len(questionIDs))
	for i, id := range questionIDs {
		values[i+1] = id
		placeholders[i] = fmt.Sprintf("$%d", i+2)
	}

	query := db.Query{
		Name: "question_set_repository.remove_questions_from_set",
		QueryRaw: fmt.Sprintf(
			"DELETE FROM question_set WHERE set_id = $1 AND question_id IN (%s)",
			strings.Join(placeholders, ", "),
		),
	}

	_, err := r.db.DB().ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}
