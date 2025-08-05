package pg

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/category/pg/converter"
)

func (r *repo) Create(ctx context.Context, category *model.NewCategory) (int64, error) {
	categoryRepo := converter.NewCategoryToPGSQL(category)
	query := db.Query{
		Name:     "category_repository.create",
		QueryRaw: `INSERT INTO category(name, domain_id) VALUES ($1, $2) RETURNING id;`,
	}

	var id int64
	err := r.db.DB().ScanOne(ctx, &id, query, categoryRepo.Name, category.DomainID)
	if err != nil {
		return 0, err
	}

	return id, nil
}
