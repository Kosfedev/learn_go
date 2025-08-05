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

func (r *repo) getOptions(ctx context.Context, id int64) ([]*model.QuestionOption, error) {
	questionOptionsServ := make([]*model.QuestionOption, 0)
	query := db.Query{
		Name:     "question_repository.add_options",
		QueryRaw: `SELECT * FROM question_option WHERE question_id = $1`,
	}
	rows, err := r.db.DB().QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		questionOptionRepo := &modelRepo.QuestionOption{}
		err = rows.Scan(&questionOptionRepo.ID, &questionOptionRepo.QuestionID, &questionOptionRepo.Text, &questionOptionRepo.IsCorrect)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}

			return nil, err
		}
		questionOptionsServ = append(questionOptionsServ, converter.QuestionOptionFromPGSQL(questionOptionRepo))
	}

	return questionOptionsServ, nil
}
