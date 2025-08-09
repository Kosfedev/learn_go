package serviceprovider

import (
	"context"

	"github.com/Kosfedev/learn_go/internal/repository"
	categoryPGRepository "github.com/Kosfedev/learn_go/internal/repository/category/pg"
	domainPGRepository "github.com/Kosfedev/learn_go/internal/repository/domain/pg"
	questionPGRepository "github.com/Kosfedev/learn_go/internal/repository/question/pg"
	questionFormPGRepository "github.com/Kosfedev/learn_go/internal/repository/question_form/pg"
	questionOptionPGRepository "github.com/Kosfedev/learn_go/internal/repository/question_option/pg"
	questionSetPGRepository "github.com/Kosfedev/learn_go/internal/repository/question_set/pg"
	questionSubcategoryPGRepository "github.com/Kosfedev/learn_go/internal/repository/question_subcategory/pg"
	setPGRepository "github.com/Kosfedev/learn_go/internal/repository/set/pg"
	subcategoryPGRepository "github.com/Kosfedev/learn_go/internal/repository/subcategory/pg"
)

func (sp *ServiceProvider) QuestionRepository(ctx context.Context) repository.QuestionRepository {
	if sp.questionRepo == nil {
		sp.questionRepo = questionPGRepository.NewRepository(sp.DBClient(ctx))
	}

	return sp.questionRepo
}

func (sp *ServiceProvider) QuestionOptionRepository(ctx context.Context) repository.QuestionOptionRepository {
	if sp.questionOptionRepo == nil {
		sp.questionOptionRepo = questionOptionPGRepository.NewRepository(sp.DBClient(ctx))
	}

	return sp.questionOptionRepo
}

func (sp *ServiceProvider) SetRepository(ctx context.Context) repository.SetRepository {
	if sp.setRepo == nil {
		sp.setRepo = setPGRepository.NewRepository(sp.DBClient(ctx))
	}

	return sp.setRepo
}

func (sp *ServiceProvider) QuestionSubcategoryRepository(ctx context.Context) repository.QuestionSubcategoryRepository {
	if sp.questionSubcategoryRepo == nil {
		sp.questionSubcategoryRepo = questionSubcategoryPGRepository.NewRepository(sp.DBClient(ctx))
	}

	return sp.questionSubcategoryRepo
}

func (sp *ServiceProvider) QuestionSetRepository(ctx context.Context) repository.QuestionSetRepository {
	if sp.questionSetRepo == nil {
		sp.questionSetRepo = questionSetPGRepository.NewRepository(sp.DBClient(ctx))
	}

	return sp.questionSetRepo
}

func (sp *ServiceProvider) QuestionFormRepository(ctx context.Context) repository.QuestionFormRepository {
	if sp.questionFormRepo == nil {
		sp.questionFormRepo = questionFormPGRepository.NewRepository(sp.DBClient(ctx))
	}

	return sp.questionFormRepo
}

func (sp *ServiceProvider) DomainRepository(ctx context.Context) repository.DomainRepository {
	if sp.domainRepo == nil {
		sp.domainRepo = domainPGRepository.NewRepository(sp.DBClient(ctx))
	}

	return sp.domainRepo
}

func (sp *ServiceProvider) CategoryRepository(ctx context.Context) repository.CategoryRepository {
	if sp.categoryRepo == nil {
		sp.categoryRepo = categoryPGRepository.NewRepository(sp.DBClient(ctx))
	}

	return sp.categoryRepo
}

func (sp *ServiceProvider) SubcategoryRepository(ctx context.Context) repository.SubcategoryRepository {
	if sp.subcategoryRepo == nil {
		sp.subcategoryRepo = subcategoryPGRepository.NewRepository(sp.DBClient(ctx))
	}

	return sp.subcategoryRepo
}
