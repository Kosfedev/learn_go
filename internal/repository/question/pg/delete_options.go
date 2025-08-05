package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/client/db"
)

func (r *repo) DeleteOptions(ctx context.Context, ids []int64) error {
	if len(ids) == 0 {
		return nil
	}

	idsRepo := make([]interface{}, len(ids))
	placeholders := make([]string, len(ids))
	for i, id := range ids {
		idsRepo[i] = id
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	query := db.Query{
		Name: "question_repository.delete_options",
		QueryRaw: fmt.Sprintf(
			"DELETE FROM question_option WHERE id IN (%s)",
			strings.Join(placeholders, ", "),
		),
	}
	_, err := r.db.DB().ExecContext(ctx, query, idsRepo...)
	if err != nil {
		return err
	}

	return nil
}
