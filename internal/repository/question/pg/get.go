package pg

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/question/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/question/pg/model"
)

func (r *repo) Get(ctx context.Context, id int64) (*model.Question, error) {
	questionRepo := &modelRepo.Question{}
	query := db.Query{
		Name:     "question_repository.get",
		QueryRaw: `SELECT * FROM question WHERE id = $1;`,
	}
	err := r.db.DB().ScanOne(ctx, questionRepo, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	questionServ := converter.QuestionFromPGSQL(questionRepo)
	questionOptionsServ, err := r.getOptions(ctx, id)
	if err != nil {
		return nil, err
	}

	questionServ.Options = questionOptionsServ

	return questionServ, nil
}
