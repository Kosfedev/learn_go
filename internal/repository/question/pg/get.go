package pg

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/question/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/question/pg/model"
)

func (r *repo) Get(ctx context.Context, id int) (*model.Question, error) {
	questionRepo := &modelRepo.Question{}
	query := `SELECT * FROM question WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&questionRepo.Id, &questionRepo.Text, &questionRepo.Type, &questionRepo.ReferenceAnswer, &questionRepo.CreatedAt, &questionRepo.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	questionServ := converter.QuestionFromPGSQL(questionRepo)

	questionOptionsServ := make([]*model.QuestionOption, 0)
	query = `SELECT * FROM question_option WHERE question_id = $1`
	rows, err := r.db.Query(query, id)

	for rows.Next() {
		questionOptionRepo := &modelRepo.QuestionOption{}
		err = rows.Scan(&questionOptionRepo.Id, &questionOptionRepo.QuestionId, &questionOptionRepo.Text, &questionOptionRepo.IsCorrect)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}

			return nil, err
		}
		questionOptionsServ = append(questionOptionsServ, converter.QuestionOptionFromPGSQL(questionOptionRepo))
	}

	questionServ.Options = questionOptionsServ

	return questionServ, nil
}
