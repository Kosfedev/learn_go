package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/question/pg/converter"
)

func (r *repo) Update(ctx context.Context, id int, updatedQuestion *model.UpdatedQuestion) error {
	updatedQuestionRepo := converter.UpdatedQuestionToPGSQL(updatedQuestion)
	index := 2
	values := []interface{}{id}
	query := "UPDATE question SET"

	if updatedQuestionRepo.Text.Valid {
		values = append(values, updatedQuestionRepo.Text)
		query += fmt.Sprintf(" text = $%d,", index)
		index++
	}

	if updatedQuestionRepo.Type.Valid {
		values = append(values, updatedQuestionRepo.Type)
		query += fmt.Sprintf(" type = $%d,", index)
		index++
	}

	if updatedQuestionRepo.ReferenceAnswer.Valid {
		values = append(values, updatedQuestionRepo.ReferenceAnswer)
		query += fmt.Sprintf(" reference_answer = $%d,", index)
		index++
	}

	query = strings.TrimSuffix(query, ",")
	query += " WHERE id = $1"

	_, err := r.db.ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}
