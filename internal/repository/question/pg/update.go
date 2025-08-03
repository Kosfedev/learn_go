package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/question/pg/converter"
)

func (r *repo) Update(ctx context.Context, id int, updatedQuestion *model.UpdatedQuestion) error {
	updatedQuestionRepo := converter.UpdatedQuestionToPGSQL(updatedQuestion)
	index := 2
	values := []interface{}{id}
	queryRaw := "UPDATE question SET"

	if updatedQuestionRepo.Text.Valid {
		values = append(values, updatedQuestionRepo.Text)
		queryRaw += fmt.Sprintf(" text = $%d,", index)
		index++
	}

	if updatedQuestionRepo.Type.Valid {
		values = append(values, updatedQuestionRepo.Type)
		queryRaw += fmt.Sprintf(" type = $%d,", index)
		index++
	}

	if updatedQuestionRepo.ReferenceAnswer.Valid {
		values = append(values, updatedQuestionRepo.ReferenceAnswer)
		queryRaw += fmt.Sprintf(" reference_answer = $%d,", index)
	}

	queryRaw = strings.TrimSuffix(queryRaw, ",")
	queryRaw += " WHERE id = $1"
	query := db.Query{
		Name:     "question_repository.update",
		QueryRaw: queryRaw,
	}

	_, err := r.db.DB().ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}
