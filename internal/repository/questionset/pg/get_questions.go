package pg

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Kosfedev/learn_go/internal/client/db"
	modelRepo "github.com/Kosfedev/learn_go/internal/repository/questionset/pg/model"
)

func (r *repo) getQuestions(ctx context.Context, id int) ([]int, error) {
	questionSetQuestionsServ := make([]int, 0)
	query := db.Query{
		Name:     "question_set_repository.get_questions",
		QueryRaw: `SELECT * FROM question_set_question WHERE question_set_id = $1;`,
	}
	rows, err := r.db.DB().QueryContext(ctx, query, id)
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

	return questionSetQuestionsServ, nil
}
