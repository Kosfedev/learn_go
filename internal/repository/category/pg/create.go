package pg

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/category/pg/converter"
)

func (r *repo) Create(ctx context.Context, category *model.NewCategory) (int, error) {
	categoryRepo := converter.NewCategoryToPGSQL(category)
	query := `INSERT INTO category(name, domain_id) VALUES ($1, $2) RETURNING id`

	var id int
	err := r.db.QueryRowContext(ctx, query, categoryRepo.Name, category.DomainId).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
