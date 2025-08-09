package serviceprovider

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/service"
	categoryService "github.com/Kosfedev/learn_go/internal/service/category"
	domainService "github.com/Kosfedev/learn_go/internal/service/domain"
	questionService "github.com/Kosfedev/learn_go/internal/service/question"
	questionFormService "github.com/Kosfedev/learn_go/internal/service/questionform"
	questionFormUpdaterService "github.com/Kosfedev/learn_go/internal/service/questionformupdater"
	questionSetService "github.com/Kosfedev/learn_go/internal/service/set"
	subcategoryService "github.com/Kosfedev/learn_go/internal/service/subcategory"
)

func (sp *ServiceProvider) QuestionService(ctx context.Context) service.QuestionService {
	if sp.questionServ == nil {
		sp.questionServ = questionService.NewService(
			sp.QuestionRepository(ctx),
			sp.QuestionOptionRepository(ctx),
			sp.QuestionSubcategoryRepository(ctx),
			sp.SubcategoryRepository(ctx),
		)
	}

	return sp.questionServ
}

func (sp *ServiceProvider) QuestionSetService(ctx context.Context) service.SetService {
	if sp.setServ == nil {
		sp.setServ = questionSetService.NewService(sp.SetRepository(ctx))
	}

	return sp.setServ
}

func (sp *ServiceProvider) DomainService(ctx context.Context) service.DomainService {
	if sp.domainServ == nil {
		sp.domainServ = domainService.NewService(sp.DomainRepository(ctx))
	}

	return sp.domainServ
}

func (sp *ServiceProvider) CategoryService(ctx context.Context) service.CategoryService {
	if sp.categoryServ == nil {
		sp.categoryServ = categoryService.NewService(sp.CategoryRepository(ctx))
	}

	return sp.categoryServ
}

func (sp *ServiceProvider) SubcategoryService(ctx context.Context) service.SubcategoryService {
	if sp.subcategoryServ == nil {
		sp.subcategoryServ = subcategoryService.NewService(sp.SubcategoryRepository(ctx))
	}

	return sp.subcategoryServ
}

func (sp *ServiceProvider) QuestionFormService(ctx context.Context) service.QuestionFormService {
	if sp.questionFormServ == nil {
		sp.questionFormServ = questionFormService.NewService(
			sp.QuestionFormRepository(ctx),
		)
	}

	return sp.questionFormServ
}

func (sp *ServiceProvider) QuestionFormUpdaterService(ctx context.Context) service.QuestionFormUpdaterService {
	if sp.questionFormUpdaterServ == nil {
		sp.questionFormUpdaterServ = questionFormUpdaterService.NewService(
			sp.QuestionRepository(ctx),
			sp.QuestionOptionRepository(ctx),
		)
	}

	return sp.questionFormUpdaterServ
}
