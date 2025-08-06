package pg

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/set/pg/converter"
)

func (r *repo) Create(ctx context.Context, newSet *model.NewSet) (int64, error) {
	newSetRepo := converter.NewSetToPGSQL(newSet)
	query := db.Query{
		Name:     "set_repository.create",
		QueryRaw: `INSERT INTO set (name) VALUES ($1) RETURNING id;`,
	}

	var setID int64
	err := r.db.DB().ScanOne(ctx, &setID, query, newSetRepo.Name)
	if err != nil {
		return 0, err
	}

	// TODO: нужна транзакция
	err = r.addQuestions(ctx, setID, newSet.QuestionIDs)
	if err != nil {
		return 0, err
	}

	return setID, nil
}
