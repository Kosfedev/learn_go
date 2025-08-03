package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/category/pg/converter"
)

func (r *repo) Update(ctx context.Context, id int, category *model.UpdatedCategory) error {
	categoryRepo := converter.UpdatedCategoryToPGSQL(category)
	values := []interface{}{id}
	queryRaw := "UPDATE category SET"
	index := 2

	if categoryRepo.Name.Valid {
		values = append(values, categoryRepo.Name)
		queryRaw += fmt.Sprintf(" name = $%d,", index)
		index++
	}

	if categoryRepo.DomainId.Valid {
		values = append(values, categoryRepo.DomainId)
		queryRaw += fmt.Sprintf(" domain_id = $%d,", index)
	}

	queryRaw = strings.TrimSuffix(queryRaw, ",")
	queryRaw += " WHERE id = $1"
	query := db.Query{
		Name:     "category_update",
		QueryRaw: queryRaw,
	}
	_, err := r.db.DB().ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}
