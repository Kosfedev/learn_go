package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/question/pg/converter"
)

func (r *repo) Create(_ context.Context, newQuestion *model.NewQuestion) (int, error) {
	newQuestionRepo := converter.NewQuestionToPGSQL(newQuestion)
	query := `
		INSERT INTO question (text, type, reference_answer) 
		VALUES ($1, $2, $3)
		RETURNING id;`

	var questionId int
	err := r.db.QueryRow(query, newQuestionRepo.Text, newQuestionRepo.Type, newQuestionRepo.ReferenceAnswer).Scan(&questionId)
	if err != nil {
		return 0, err
	}

	// TODO: вынести в общую функцию?
	// TODO: нужна транзакция
	if len(newQuestion.Options) > 0 {
		var values []interface{}
		nArgs := 3
		query = `INSERT INTO question_option (question_id, text, is_correct) VALUES`

		for i, option := range newQuestion.Options {
			query += fmt.Sprintf(" ($%d, $%d, $%d),", i*nArgs+1, i*nArgs+2, i*nArgs+3)
			newOptionRepo := converter.NewQuestionOptionToPGSQL(questionId, option)
			values = append(values, newOptionRepo.QuestionId, newOptionRepo.Text, newOptionRepo.IsCorrect)
		}

		query = strings.TrimSuffix(query, ",")

		_, err = r.db.Query(query, values...)
		if err != nil {
			return 0, err
		}
	}

	return questionId, nil
}
