package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/category/pg/converter"
)

func (r *repo) Update(ctx context.Context, id int, category *model.UpdatedCategory) error {
	categoryRepo := converter.UpdatedCategoryToPGSQL(category)
	values := []interface{}{id}
	query := "UPDATE category SET"
	index := 2

	if categoryRepo.Name.Valid {
		values = append(values, categoryRepo.Name)
		query += fmt.Sprintf(" name = $%d,", index)
		index++
	}

	if categoryRepo.DomainId.Valid {
		values = append(values, categoryRepo.DomainId)
		query += fmt.Sprintf(" domain_id = $%d,", index)
	}

	query = strings.TrimSuffix(query, ",")
	query += " WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}
