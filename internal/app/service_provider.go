package app

import (
	"context"

	categoryImplementation "github.com/Kosfedev/learn_go/internal/api/category"
	domainImplementation "github.com/Kosfedev/learn_go/internal/api/domain"
	quesionImplementation "github.com/Kosfedev/learn_go/internal/api/question"
	subcategoryImplementation "github.com/Kosfedev/learn_go/internal/api/subcategory"
	"github.com/Kosfedev/learn_go/internal/service"
	categoryService "github.com/Kosfedev/learn_go/internal/service/category"
	domainService "github.com/Kosfedev/learn_go/internal/service/domain"
	questionService "github.com/Kosfedev/learn_go/internal/service/question"
	subcategoryService "github.com/Kosfedev/learn_go/internal/service/subcategory"
)

type serviceProvider struct {
	questionServ    service.QuestionService
	domainServ      service.DomainService
	categoryServ    service.CategoryService
	subcategoryServ service.SubcategoryService

	questionImpl    *quesionImplementation.Implementation
	domainImpl      *domainImplementation.Implementation
	categoryImpl    *categoryImplementation.Implementation
	subcategoryImpl *subcategoryImplementation.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (sp *serviceProvider) QuestionService(_ context.Context) service.QuestionService {
	if sp.questionServ == nil {
		sp.questionServ = questionService.NewService()
	}

	return sp.questionServ
}

func (sp *serviceProvider) DomainService(_ context.Context) service.DomainService {
	if sp.domainServ == nil {
		sp.domainServ = domainService.NewService()
	}

	return sp.domainServ
}

func (sp *serviceProvider) CategoryService(_ context.Context) service.CategoryService {
	if sp.categoryServ == nil {
		sp.categoryServ = categoryService.NewService()
	}

	return sp.categoryServ
}

func (sp *serviceProvider) SubcategoryService(_ context.Context) service.SubcategoryService {
	if sp.subcategoryServ == nil {
		sp.subcategoryServ = subcategoryService.NewService()
	}

	return sp.subcategoryServ
}

func (sp *serviceProvider) QuestionImplementation(ctx context.Context) *quesionImplementation.Implementation {
	if sp.questionImpl == nil {
		sp.questionImpl = quesionImplementation.NewImplementation(sp.QuestionService(ctx))
	}

	return sp.questionImpl
}

func (sp *serviceProvider) DomainImplementation(ctx context.Context) *domainImplementation.Implementation {
	if sp.domainImpl == nil {
		sp.domainImpl = domainImplementation.NewImplementation(sp.DomainService(ctx))
	}

	return sp.domainImpl
}

func (sp *serviceProvider) CategoryImplementation(ctx context.Context) *categoryImplementation.Implementation {
	if sp.categoryImpl == nil {
		sp.categoryImpl = categoryImplementation.NewImplementation(sp.CategoryService(ctx))
	}

	return sp.categoryImpl
}

func (sp *serviceProvider) SubcategoryImplementation(ctx context.Context) *subcategoryImplementation.Implementation {
	if sp.subcategoryImpl == nil {
		sp.subcategoryImpl = subcategoryImplementation.NewImplementation(sp.SubcategoryService(ctx))
	}

	return sp.subcategoryImpl
}
