package pg

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/category/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/category/pg/model"
)

func (r *repo) Get(ctx context.Context, id int64) (*model.Category, error) {
	query := db.Query{
		Name:     "category_repository.get",
		QueryRaw: `SELECT * FROM category WHERE id = $1;`,
	}
	categoryRepo := &modelRepo.Category{}
	err := r.db.DB().ScanOne(ctx, categoryRepo, query, id)
	if err != nil {
		return nil, err
	}

	category := converter.CategoryFromPGSQL(categoryRepo)

	return category, nil
}
