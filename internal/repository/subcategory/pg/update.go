package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/subcategory/pg/converter"
)

func (r *repo) Update(ctx context.Context, id int, subcategory *model.UpdatedSubcategory) error {
	subcategoryRepo := converter.UpdatedSubcategoryToPGSQL(subcategory)
	values := []interface{}{id}
	query := "UPDATE subcategory SET"
	index := 2

	if subcategoryRepo.Name.Valid {
		values = append(values, subcategoryRepo.Name)
		query += fmt.Sprintf(" name = $%d,", index)
		index++
	}

	if subcategoryRepo.CategoryId.Valid {
		values = append(values, subcategoryRepo.CategoryId)
		query += fmt.Sprintf(" category_id = $%d,", index)
	}

	query = strings.TrimSuffix(query, ",")
	query += " WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}
