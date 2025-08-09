package subcategory

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/model"
)

func (s *serv) Update(ctx context.Context, id int64, updatedCategory *model.UpdatedSubcategory) error {
	err := s.repo.Update(ctx, id, updatedCategory)
	if err != nil {
		return err
	}

	return nil
}
