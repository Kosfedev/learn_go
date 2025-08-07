package pg

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/client/db"
)

func (r *repo) Delete(ctx context.Context, id int64) error {
	query := db.Query{
		Name:     "set_repository.delete",
		QueryRaw: `DELETE FROM set WHERE id = $1;`,
	}
	_, err := r.db.DB().ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
