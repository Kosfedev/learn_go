package questionformupdater

import (
	"context"

	desc "github.com/Kosfedev/learn_go/pkg/question_form_updater_v1"
)

func (impl *Implementation) RemoveSubcategories(ctx context.Context, req *desc.RemoveSubcategoriesRequest) (*desc.RemoveSubcategoriesResponse, error) {
	err := impl.questionFormUpdaterService.RemoveSubcategories(ctx, req.QuestionId, req.SubcategoryIds)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
