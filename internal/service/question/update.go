package question

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (s *serv) Update(ctx context.Context, id int64, updatedQuestion *model.UpdatedQuestion) error {
	err := s.questionRepo.Update(ctx, id, updatedQuestion)
	if err != nil {
		return err
	}

	return nil
}
