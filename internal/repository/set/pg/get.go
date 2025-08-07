package pg

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/set/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/set/pg/model"
)

func (r *repo) Get(ctx context.Context, id int64) (*model.Set, error) {
	setRepo := &modelRepo.Set{}
	query := db.Query{
		Name:     "set_repository.get",
		QueryRaw: `SELECT * FROM set WHERE id = $1`,
	}
	err := r.db.DB().ScanOne(ctx, setRepo, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	set := converter.SetFromPGSQL(setRepo)
	questionIDs, err := r.getQuestions(ctx, id)
	if err != nil {
		return nil, err
	}

	set.QuestionIDs = questionIDs

	return set, nil
}
