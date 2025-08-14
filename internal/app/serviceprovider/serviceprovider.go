package serviceprovider

import (
	"context"
	"log"

	categoryImplementation "github.com/Kosfedev/learn_go/internal/api/category"
	domainImplementation "github.com/Kosfedev/learn_go/internal/api/domain"
	quesionImplementation "github.com/Kosfedev/learn_go/internal/api/question"
	quesionFormImplementation "github.com/Kosfedev/learn_go/internal/api/questionform"
	quesionFormUpdaterImplementation "github.com/Kosfedev/learn_go/internal/api/questionformupdater"
	setImplementation "github.com/Kosfedev/learn_go/internal/api/set"
	subcategoryImplementation "github.com/Kosfedev/learn_go/internal/api/subcategory"
	"github.com/Kosfedev/learn_go/internal/client/db"
	"github.com/Kosfedev/learn_go/internal/client/db/pg"
	"github.com/Kosfedev/learn_go/internal/client/db/transactor"
	"github.com/Kosfedev/learn_go/internal/config"
	"github.com/Kosfedev/learn_go/internal/config/env"
	"github.com/Kosfedev/learn_go/internal/repository"
	"github.com/Kosfedev/learn_go/internal/service"
)

type ServiceProvider struct {
	pgConfig     config.PGConfig
	grpcConfig   config.GRPCConfig
	grpcGWConfig config.GRPCGWConfig

	dbClient  db.Client
	txManager db.TxManager

	questionRepo            repository.QuestionRepository
	questionSetRepo         repository.QuestionSetRepository
	questionSubcategoryRepo repository.QuestionSubcategoryRepository
	questionFormRepo        repository.QuestionFormRepository
	questionOptionRepo      repository.QuestionOptionRepository
	setRepo                 repository.SetRepository
	setFormRepo             repository.SetFormRepository
	domainRepo              repository.DomainRepository
	categoryRepo            repository.CategoryRepository
	subcategoryRepo         repository.SubcategoryRepository

	questionServ            service.QuestionService
	questionFormServ        service.QuestionFormService
	questionFormUpdaterServ service.QuestionFormUpdaterService
	setServ                 service.SetService
	setFormServ             service.SetFormService
	domainServ              service.DomainService
	categoryServ            service.CategoryService
	subcategoryServ         service.SubcategoryService

	questionImpl            *quesionImplementation.Implementation
	questionFormImpl        *quesionFormImplementation.Implementation
	questionFormUpdaterImpl *quesionFormUpdaterImplementation.Implementation
	setImpl                 *setImplementation.Implementation
	domainImpl              *domainImplementation.Implementation
	categoryImpl            *categoryImplementation.Implementation
	subcategoryImpl         *subcategoryImplementation.Implementation
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

func (sp *ServiceProvider) PGConfig() config.PGConfig {
	if sp.pgConfig == nil {
		cfg, err := env.NewPgConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		sp.pgConfig = cfg
	}

	return sp.pgConfig
}

func (sp *ServiceProvider) GRPCConfig() config.GRPCConfig {
	if sp.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		sp.grpcConfig = cfg
	}

	return sp.grpcConfig
}

func (sp *ServiceProvider) GRPCGWConfig() config.GRPCGWConfig {
	if sp.grpcGWConfig == nil {
		cfg, err := env.NewGRPCGWConfig()
		if err != nil {
			log.Fatalf("failed to get grpc-gw config: %s", err.Error())
		}

		sp.grpcGWConfig = cfg
	}

	return sp.grpcGWConfig
}

func (sp *ServiceProvider) DBClient(ctx context.Context) db.Client {
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

func (sp *ServiceProvider) TxManager(ctx context.Context) db.TxManager {
	if sp.txManager == nil {
		sp.txManager = transactor.NewManager(sp.DBClient(ctx).DB())
	}

	return sp.txManager
}
