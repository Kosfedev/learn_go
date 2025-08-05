package pg

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/questionset/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/questionset/pg/model"
)

func (r *repo) Get(ctx context.Context, id int64) (*model.QuestionSet, error) {
	questionSetRepo := &modelRepo.QuestionSet{}
	query := db.Query{
		Name:     "question_set_repository.get",
		QueryRaw: `SELECT * FROM question_set WHERE id = $1`,
	}
	err := r.db.DB().ScanOne(ctx, questionSetRepo, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	questionSetServ := converter.QuestionSetFromPGSQL(questionSetRepo)
	questionSetQuestionsServ, err := r.getQuestions(ctx, id)
	if err != nil {
		return nil, err
	}

	questionSetServ.QuestionIDs = questionSetQuestionsServ

	return questionSetServ, nil
}
