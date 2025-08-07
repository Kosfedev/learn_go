package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/set/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/set/pg/model"
)

func (r *repo) ListByIDs(ctx context.Context, ids []int64) ([]*model.Set, error) {
	idsRepo := make([]interface{}, len(ids))
	placeholders := make([]string, len(ids))
	for i, id := range ids {
		idsRepo[i] = id
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	query := db.Query{
		Name: "set_repository.list_by_ids",
		QueryRaw: fmt.Sprintf(
			"SELECT * FROM set WHERE id IN (%s)",
			strings.Join(placeholders, ", "),
		),
	}
	setsRepo := make([]*modelRepo.Set, len(ids))
	err := r.db.DB().ScanAll(ctx, &setsRepo, query, idsRepo...)
	if err != nil {
		return nil, err
	}

	set := converter.SetsFromPGSQL(setsRepo)

	return set, nil
}
