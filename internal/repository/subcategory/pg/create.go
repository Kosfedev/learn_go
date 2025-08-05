package pg

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/subcategory/pg/converter"
)

func (r *repo) Create(ctx context.Context, subcategory *model.NewSubcategory) (int64, error) {
	subcategoryRepo := converter.NewSubcategoryToPGSQL(subcategory)
	query := db.Query{
		Name:     "subcategory_repository.create",
		QueryRaw: `INSERT INTO subcategory(name, category_id) VALUES ($1, $2) RETURNING id;`,
	}

	var id int64
	err := r.db.DB().ScanOne(ctx, &id, query, subcategoryRepo.Name, subcategory.CategoryID)
	if err != nil {
		return 0, err
	}

	return id, nil
}
