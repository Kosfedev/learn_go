package pg

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/category/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/category/pg/model"
)

func (r *repo) Get(ctx context.Context, id int) (*model.Category, error) {
	query := `SELECT * FROM category WHERE id = $1`
	categoryRepo := &modelRepo.Category{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&categoryRepo.Id, &categoryRepo.Name, &categoryRepo.DomainId, &categoryRepo.CreatedAt, &categoryRepo.UpdatedAt)
	if err != nil {
		return nil, err
	}

	category := converter.CategoryFromPGSQL(categoryRepo)

	return category, nil
}
