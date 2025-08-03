package pg

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/subcategory/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/subcategory/pg/model"
)

func (r *repo) Get(ctx context.Context, id int) (*model.Subcategory, error) {
	query := db.Query{
		Name:     "subcategory_repository.get",
		QueryRaw: `SELECT * FROM subcategory WHERE id = $1;`,
	}
	subcategoryRepo := &modelRepo.Subcategory{}
	err := r.db.DB().ScanOne(ctx, subcategoryRepo, query, id)
	if err != nil {
		return nil, err
	}

	subcategory := converter.SubcategoryFromPGSQL(subcategoryRepo)

	return subcategory, nil
}
