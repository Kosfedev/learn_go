package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/question/pg/converter"
)

func (r *repo) AddOptions(ctx context.Context, questionId int, options []*model.NewQuestionOption) error {
	if len(options) == 0 {
		return nil
	}

	// TODO: вынести в общую функцию?
	var values []interface{}
	nArgs := 3
	query := `INSERT INTO question_option (question_id, text, is_correct) VALUES`

	for i, option := range options {
		query += fmt.Sprintf(" ($%d, $%d, $%d),", i*nArgs+1, i*nArgs+2, i*nArgs+3)
		newOptionRepo := converter.NewQuestionOptionToPGSQL(questionId, option)
		values = append(values, newOptionRepo.QuestionId, newOptionRepo.Text, newOptionRepo.IsCorrect)
	}

	query = strings.TrimSuffix(query, ",")

	_, err := r.db.QueryContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}
