package pg

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/subcategory/pg/converter"
)

func (r *repo) Create(ctx context.Context, subcategory *model.NewSubcategory) (int, error) {
	subcategoryRepo := converter.NewSubcategoryToPGSQL(subcategory)
	query := `INSERT INTO subcategory(name, category_id) VALUES ($1, $2) RETURNING id`

	var id int
	err := r.db.QueryRowContext(ctx, query, subcategoryRepo.Name, subcategory.CategoryId).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
