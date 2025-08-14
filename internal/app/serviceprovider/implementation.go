package serviceprovider

import (
	"context"

	categoryImplementation "github.com/Kosfedev/learn_go/internal/api/category"
	domainImplementation "github.com/Kosfedev/learn_go/internal/api/domain"
	quesionImplementation "github.com/Kosfedev/learn_go/internal/api/question"
	quesionFormImplementation "github.com/Kosfedev/learn_go/internal/api/questionform"
	quesionFormUpdaterImplementation "github.com/Kosfedev/learn_go/internal/api/questionformupdater"
	setImplementation "github.com/Kosfedev/learn_go/internal/api/set"
	setFormImplementation "github.com/Kosfedev/learn_go/internal/api/setform"
	subcategoryImplementation "github.com/Kosfedev/learn_go/internal/api/subcategory"
)

func (sp *ServiceProvider) QuestionImplementation(ctx context.Context) *quesionImplementation.Implementation {
	if sp.questionImpl == nil {
		sp.questionImpl = quesionImplementation.NewImplementation(sp.QuestionService(ctx))
	}

	return sp.questionImpl
}

func (sp *ServiceProvider) QuestionFormImplementation(ctx context.Context) *quesionFormImplementation.Implementation {
	if sp.questionFormImpl == nil {
		sp.questionFormImpl = quesionFormImplementation.NewImplementation(sp.QuestionFormService(ctx))
	}

	return sp.questionFormImpl
}

func (sp *ServiceProvider) QuestionFormUpdaterImplementation(ctx context.Context) *quesionFormUpdaterImplementation.Implementation {
	if sp.questionFormUpdaterImpl == nil {
		sp.questionFormUpdaterImpl = quesionFormUpdaterImplementation.NewImplementation(sp.QuestionFormUpdaterService(ctx))
	}

	return sp.questionFormUpdaterImpl
}

func (sp *ServiceProvider) SetImplementation(ctx context.Context) *setImplementation.Implementation {
	if sp.setImpl == nil {
		sp.setImpl = setImplementation.NewImplementation(sp.SetService(ctx))
	}

	return sp.setImpl
}

func (sp *ServiceProvider) SetFormImplementation(ctx context.Context) *setFormImplementation.Implementation {
	if sp.setFormImpl == nil {
		sp.setFormImpl = setFormImplementation.NewImplementation(sp.SetFormService(ctx))
	}

	return sp.setFormImpl
}

func (sp *ServiceProvider) DomainImplementation(ctx context.Context) *domainImplementation.Implementation {
	if sp.domainImpl == nil {
		sp.domainImpl = domainImplementation.NewImplementation(sp.DomainService(ctx))
	}

	return sp.domainImpl
}

func (sp *ServiceProvider) CategoryImplementation(ctx context.Context) *categoryImplementation.Implementation {
	if sp.categoryImpl == nil {
		sp.categoryImpl = categoryImplementation.NewImplementation(sp.CategoryService(ctx))
	}

	return sp.categoryImpl
}

func (sp *ServiceProvider) SubcategoryImplementation(ctx context.Context) *subcategoryImplementation.Implementation {
	if sp.subcategoryImpl == nil {
		sp.subcategoryImpl = subcategoryImplementation.NewImplementation(sp.SubcategoryService(ctx))
	}

	return sp.subcategoryImpl
}
