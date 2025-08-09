package questionformupdater

import (
	"context"

	desc "github.com/Kosfedev/learn_go/pkg/question_form_updater_v1"
)

func (impl *Implementation) AddSubcategories(ctx context.Context, req *desc.AddSubcategoriesRequest) (*desc.AddSubcategoriesResponse, error) {
	err := impl.questionFormUpdaterService.AddSubcategories(ctx, req.QuestionId, req.SubcategoryIds)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
