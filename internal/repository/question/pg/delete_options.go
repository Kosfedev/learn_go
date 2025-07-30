package pg

import (
	"context"
	"fmt"
	"strings"
)

func (r *repo) DeleteOptions(_ context.Context, ids []int) error {
	if len(ids) == 0 {
		return nil
	}

	idsRepo := make([]interface{}, len(ids))
	placeholders := make([]string, len(ids))
	for i, id := range ids {
		idsRepo[i] = int32(id)
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	query := fmt.Sprintf(
		"DELETE FROM question_option WHERE id IN (%s)",
		strings.Join(placeholders, ", "),
	)
	_, err := r.db.Exec(query, idsRepo...)
	if err != nil {
		return err
	}

	return nil
}
