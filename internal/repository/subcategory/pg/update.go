package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/subcategory/pg/converter"
)

func (r *repo) Update(ctx context.Context, id int, subcategory *model.UpdatedSubcategory) error {
	subcategoryRepo := converter.UpdatedSubcategoryToPGSQL(subcategory)
	values := []interface{}{id}
	queryRaw := "UPDATE subcategory SET"
	index := 2

	if subcategoryRepo.Name.Valid {
		values = append(values, subcategoryRepo.Name)
		queryRaw += fmt.Sprintf(" name = $%d,", index)
		index++
	}

	if subcategoryRepo.CategoryId.Valid {
		values = append(values, subcategoryRepo.CategoryId)
		queryRaw += fmt.Sprintf(" category_id = $%d,", index)
	}

	queryRaw = strings.TrimSuffix(queryRaw, ",")
	queryRaw += " WHERE id = $1"
	query := db.Query{
		Name:     "subcategory_update",
		QueryRaw: queryRaw,
	}
	_, err := r.db.DB().ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}
