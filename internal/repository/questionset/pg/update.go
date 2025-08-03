package pg

import (
	"context"
	"fmt"
	"strings"

	"github.com/Kosfedev/learn_go/internal/model"
	"github.com/Kosfedev/learn_go/internal/repository/questionset/pg/converter"
)

func (r *repo) Update(ctx context.Context, id int, updatedQuestionSet *model.UpdatedQuestionSet) error {
	updatedQuestionSetRepo := converter.UpdatedQuestionSetToPGSQL(updatedQuestionSet)
	index := 2
	values := []interface{}{id}
	query := "UPDATE question_set SET"

	if updatedQuestionSetRepo.Name.Valid {
		values = append(values, updatedQuestionSetRepo.Name)
		query += fmt.Sprintf(" name = $%d,", index)
	}

	query = strings.TrimSuffix(query, ",")
	query += " WHERE id = $1"

	_, err := r.db.ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}

	return nil
}
