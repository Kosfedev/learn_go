package app

import (
	"context"
	"log"

	categoryImplementation "github.com/Kosfedev/learn_go/internal/api/category"
	domainImplementation "github.com/Kosfedev/learn_go/internal/api/domain"
	quesionImplementation "github.com/Kosfedev/learn_go/internal/api/question"
	quesionSetImplementation "github.com/Kosfedev/learn_go/internal/api/questionset"
	subcategoryImplementation "github.com/Kosfedev/learn_go/internal/api/subcategory"
	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/client/db/pg"
	"github.com/Kosfedev/learn_go/internal/config"
	"github.com/Kosfedev/learn_go/internal/config/env"
	"github.com/Kosfedev/learn_go/internal/repository"
	categoryPGRepository "github.com/Kosfedev/learn_go/internal/repository/category/pg"
	domainPGRepository "github.com/Kosfedev/learn_go/internal/repository/domain/pg"
	questionPGRepository "github.com/Kosfedev/learn_go/internal/repository/question/pg"
	questionSubcategoryPGRepository "github.com/Kosfedev/learn_go/internal/repository/question_subcategory/pg"
	questionSetPGRepository "github.com/Kosfedev/learn_go/internal/repository/questionset/pg"
	subcategoryPGRepository "github.com/Kosfedev/learn_go/internal/repository/subcategory/pg"
	"github.com/Kosfedev/learn_go/internal/service"
	categoryService "github.com/Kosfedev/learn_go/internal/service/category"
	domainService "github.com/Kosfedev/learn_go/internal/service/domain"
	questionService "github.com/Kosfedev/learn_go/internal/service/question"
	questionSetService "github.com/Kosfedev/learn_go/internal/service/questionset"
	subcategoryService "github.com/Kosfedev/learn_go/internal/service/subcategory"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient db.Client

	questionRepo            repository.QuestionRepository
	questionSetRepo         repository.QuestionSetRepository
	questionSubcategoryRepo repository.QuestionSubcategoryRepository
	domainRepo              repository.DomainRepository
	categoryRepo            repository.CategoryRepository
	subcategoryRepo         repository.SubcategoryRepository

	questionServ    service.QuestionService
	questionSetServ service.QuestionSetService
	domainServ      service.DomainService
	categoryServ    service.CategoryService
	subcategoryServ service.SubcategoryService

	questionImpl    *quesionImplementation.Implementation
	questionSetImpl *quesionSetImplementation.Implementation
	domainImpl      *domainImplementation.Implementation
	categoryImpl    *categoryImplementation.Implementation
	subcategoryImpl *subcategoryImplementation.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (sp *serviceProvider) PGConfig() config.PGConfig {
	if sp.pgConfig == nil {
		cfg, err := env.NewPgConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		sp.pgConfig = cfg
	}

	return sp.pgConfig
}

func (sp *serviceProvider) GRPCConfig() config.GRPCConfig {
	if sp.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		sp.grpcConfig = cfg
	}

	return sp.grpcConfig
}

func (sp *serviceProvider) DBClient(ctx context.Context) db.Client {
	if sp.dbClient == nil {
		cl, err := pg.New(ctx, sp.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}

		sp.dbClient = cl

		// TODO: add closer!!!!
	}

	return sp.dbClient
}

func (sp *serviceProvider) QuestionRepository(ctx context.Context) repository.QuestionRepository {
	if sp.questionRepo == nil {
		sp.questionRepo = questionPGRepository.NewRepository(sp.DBClient(ctx))
	}

	return sp.questionRepo
}

func (sp *serviceProvider) QuestionSetRepository(ctx context.Context) repository.QuestionSetRepository {
	if sp.questionSetRepo == nil {
		sp.questionSetRepo = questionSetPGRepository.NewRepository(sp.DBClient(ctx))
	}

	return sp.questionSetRepo
}

func (sp *serviceProvider) QuestionSubcategoryRepository(ctx context.Context) repository.QuestionSubcategoryRepository {
	if sp.questionSubcategoryRepo == nil {
		sp.questionSubcategoryRepo = questionSubcategoryPGRepository.NewRepository(sp.DBClient(ctx))
	}

	return sp.questionSubcategoryRepo
}

func (sp *serviceProvider) DomainRepository(ctx context.Context) repository.DomainRepository {
	if sp.domainRepo == nil {
		sp.domainRepo = domainPGRepository.NewRepository(sp.DBClient(ctx))
	}

	return sp.domainRepo
}

func (sp *serviceProvider) CategoryRepository(ctx context.Context) repository.CategoryRepository {
	if sp.categoryRepo == nil {
		sp.categoryRepo = categoryPGRepository.NewRepository(sp.DBClient(ctx))
	}

	return sp.categoryRepo
}

func (sp *serviceProvider) SubcategoryRepository(ctx context.Context) repository.SubcategoryRepository {
	if sp.subcategoryRepo == nil {
		sp.subcategoryRepo = subcategoryPGRepository.NewRepository(sp.DBClient(ctx))
	}

	return sp.subcategoryRepo
}

func (sp *serviceProvider) QuestionService(ctx context.Context) service.QuestionService {
	if sp.questionServ == nil {
		sp.questionServ = questionService.NewService(sp.QuestionRepository(ctx), sp.QuestionSubcategoryRepository(ctx), sp.SubcategoryRepository(ctx))
	}

	return sp.questionServ
}

func (sp *serviceProvider) QuestionSetService(ctx context.Context) service.QuestionSetService {
	if sp.questionSetServ == nil {
		sp.questionSetServ = questionSetService.NewService(sp.QuestionSetRepository(ctx))
	}

	return sp.questionSetServ
}

func (sp *serviceProvider) DomainService(ctx context.Context) service.DomainService {
	if sp.domainServ == nil {
		sp.domainServ = domainService.NewService(sp.DomainRepository(ctx))
	}

	return sp.domainServ
}

func (sp *serviceProvider) CategoryService(ctx context.Context) service.CategoryService {
	if sp.categoryServ == nil {
		sp.categoryServ = categoryService.NewService(sp.CategoryRepository(ctx))
	}

	return sp.categoryServ
}

func (sp *serviceProvider) SubcategoryService(ctx context.Context) service.SubcategoryService {
	if sp.subcategoryServ == nil {
		sp.subcategoryServ = subcategoryService.NewService(sp.SubcategoryRepository(ctx))
	}

	return sp.subcategoryServ
}

func (sp *serviceProvider) QuestionImplementation(ctx context.Context) *quesionImplementation.Implementation {
	if sp.questionImpl == nil {
		sp.questionImpl = quesionImplementation.NewImplementation(sp.QuestionService(ctx))
	}

	return sp.questionImpl
}

func (sp *serviceProvider) QuestionSetImplementation(ctx context.Context) *quesionSetImplementation.Implementation {
	if sp.questionSetImpl == nil {
		sp.questionSetImpl = quesionSetImplementation.NewImplementation(sp.QuestionSetService(ctx))
	}

	return sp.questionSetImpl
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
