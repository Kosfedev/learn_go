package pg

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Kosfedev/learn_go/internal/client/db"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/set/pg/model"
)

func (r *repo) getQuestions(ctx context.Context, id int64) ([]int64, error) {
	questionIDs := make([]int64, 0)
	query := db.Query{
		Name:     "set_repository.get_questions",
		QueryRaw: `SELECT * FROM question_set WHERE set_id = $1;`,
	}
	rows, err := r.db.DB().QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		questionSetRepo := &modelRepo.QuestionSet{}
		err = rows.Scan(&questionSetRepo.QuestionID, &questionSetRepo.SetID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}

			return nil, err
		}
		questionIDs = append(questionIDs, questionSetRepo.QuestionID)
	}

	return questionIDs, nil
}
