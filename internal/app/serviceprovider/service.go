package serviceprovider

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/service"
	categoryService "github.com/Kosfedev/learn_go/internal/service/category"
	domainService "github.com/Kosfedev/learn_go/internal/service/domain"
	questionService "github.com/Kosfedev/learn_go/internal/service/question"
	questionFormService "github.com/Kosfedev/learn_go/internal/service/questionform"
	questionFormUpdaterService "github.com/Kosfedev/learn_go/internal/service/questionformupdater"
	SetService "github.com/Kosfedev/learn_go/internal/service/set"
	SetFormService "github.com/Kosfedev/learn_go/internal/service/setform"
	subcategoryService "github.com/Kosfedev/learn_go/internal/service/subcategory"
)

func (sp *ServiceProvider) QuestionService(ctx context.Context) service.QuestionService {
	if sp.questionServ == nil {
		sp.questionServ = questionService.NewService(sp.QuestionRepository(ctx))
	}

	return sp.questionServ
}

func (sp *ServiceProvider) QuestionFormService(ctx context.Context) service.QuestionFormService {
	if sp.questionFormServ == nil {
		sp.questionFormServ = questionFormService.NewService(sp.QuestionFormRepository(ctx))
	}

	return sp.questionFormServ
}

func (sp *ServiceProvider) QuestionFormUpdaterService(ctx context.Context) service.QuestionFormUpdaterService {
	if sp.questionFormUpdaterServ == nil {
		sp.questionFormUpdaterServ = questionFormUpdaterService.NewService(
			sp.TxManager(ctx),
			sp.QuestionRepository(ctx),
			sp.QuestionOptionRepository(ctx),
			sp.QuestionSubcategoryRepository(ctx),
			sp.QuestionSetRepository(ctx),
		)
	}

	return sp.questionFormUpdaterServ
}

func (sp *ServiceProvider) SetService(ctx context.Context) service.SetService {
	if sp.setServ == nil {
		sp.setServ = SetService.NewService(sp.SetRepository(ctx))
	}

	return sp.setServ
}

func (sp *ServiceProvider) SetFormService(ctx context.Context) service.SetFormService {
	if sp.setFormServ == nil {
		sp.setFormServ = SetFormService.NewService(sp.SetFormRepository(ctx))
	}

	return sp.setFormServ
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
