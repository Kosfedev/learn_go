package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/subcategory/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/subcategory/pg/model"
)

func (r *repo) ListByIds(ctx context.Context, ids []int) ([]*model.Subcategory, error) {
	idsRepo := make([]interface{}, len(ids))
	placeholders := make([]string, len(ids))
	for i, id := range ids {
		idsRepo[i] = int32(id)
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	query := db.Query{
		Name: "subcategory_repository.list_by_ids",
		QueryRaw: fmt.Sprintf(
			"SELECT * FROM subcategory WHERE id IN (%s)",
			strings.Join(placeholders, ", "),
		),
	}
	subcategoriesRepo := make([]*modelRepo.Subcategory, len(ids))
	err := r.db.DB().ScanAll(ctx, &subcategoriesRepo, query, idsRepo...)
	if err != nil {
		return nil, err
	}

	subcategory := converter.SubcategoriesFromPGSQL(subcategoriesRepo)

	return subcategory, nil
}
