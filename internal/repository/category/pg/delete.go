package pg

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/client/db"
)

func (r *repo) Delete(ctx context.Context, id int) error {
	query := db.Query{
		Name:     "category_delete",
		QueryRaw: `DELETE FROM category WHERE id = $1;`,
	}
	_, err := r.db.DB().ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
