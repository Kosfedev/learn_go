package pg

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/questionset/pg/converter"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/questionset/pg/model"
)

func (r *repo) Get(ctx context.Context, id int) (*model.QuestionSet, error) {
	questionSetRepo := &modelRepo.QuestionSet{}
	query := `SELECT * FROM question_set WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&questionSetRepo.Id, &questionSetRepo.Name, &questionSetRepo.CreatedAt, &questionSetRepo.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	questionSetServ := converter.QuestionSetFromPGSQL(questionSetRepo)

	questionSetQuestionsServ := make([]int, 0)
	query = `SELECT * FROM question_set_question WHERE question_set_id = $1`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		questionSetOptionRepo := &modelRepo.QuestionSetQuestion{}
		err = rows.Scan(&questionSetOptionRepo.QuestionId, &questionSetOptionRepo.QuestionSetId)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}

			return nil, err
		}
		questionSetQuestionsServ = append(questionSetQuestionsServ, int(questionSetOptionRepo.QuestionId))
	}

	questionSetServ.QuestionIds = questionSetQuestionsServ

	return questionSetServ, nil
}
