package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	categoryImplementation "github.com/Kosfedev/learn_go/internal/api/category"
	domainImplementation "github.com/Kosfedev/learn_go/internal/api/domain"
	quesionImplementation "github.com/Kosfedev/learn_go/internal/api/question"
	subcategoryImplementation "github.com/Kosfedev/learn_go/internal/api/subcategory"
	"github.com/Kosfedev/learn_go/internal/config"
	"github.com/Kosfedev/learn_go/internal/config/env"
	"github.com/Kosfedev/learn_go/internal/repository"
	categoryPGRepository "github.com/Kosfedev/learn_go/internal/repository/category/pg"
	domainPGRepository "github.com/Kosfedev/learn_go/internal/repository/domain/pg"
	questionPGRepository "github.com/Kosfedev/learn_go/internal/repository/question/pg"
	subcategoryPGRepository "github.com/Kosfedev/learn_go/internal/repository/subcategory/pg"
	"github.com/Kosfedev/learn_go/internal/service"
	categoryService "github.com/Kosfedev/learn_go/internal/service/category"
	domainService "github.com/Kosfedev/learn_go/internal/service/domain"
	questionService "github.com/Kosfedev/learn_go/internal/service/question"
	subcategoryService "github.com/Kosfedev/learn_go/internal/service/subcategory"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	questionRepo    repository.QuestionRepository
	domainRepo      repository.DomainRepository
	categoryRepo    repository.CategoryRepository
	subcategoryRepo repository.SubcategoryRepository

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

func (sp *serviceProvider) QuestionRepository(_ context.Context) repository.QuestionRepository {
	// TODO: ПЕРЕНЕСТИ!
	db, err := sql.Open("postgres", sp.PGConfig().DSN())
	if err != nil {
		log.Fatal(err)
	}
	// TODO: defer db.Close()

	// TODO: ПЕРЕНЕСТИ!
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected!")

	if sp.questionRepo == nil {
		sp.questionRepo = questionPGRepository.NewRepository(db)
	}

	return sp.questionRepo
}

func (sp *serviceProvider) DomainRepository(_ context.Context) repository.DomainRepository {
	// TODO: ПЕРЕНЕСТИ!
	db, err := sql.Open("postgres", sp.PGConfig().DSN())
	if err != nil {
		log.Fatal(err)
	}
	// TODO: defer db.Close()

	// TODO: ПЕРЕНЕСТИ!
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected!")

	if sp.domainRepo == nil {
		sp.domainRepo = domainPGRepository.NewRepository(db)
	}

	return sp.domainRepo
}

func (sp *serviceProvider) CategoryRepository(_ context.Context) repository.CategoryRepository {
	// TODO: ПЕРЕНЕСТИ!
	db, err := sql.Open("postgres", sp.PGConfig().DSN())
	if err != nil {
		log.Fatal(err)
	}
	// TODO: defer db.Close()

	// TODO: ПЕРЕНЕСТИ!
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected!")

	if sp.categoryRepo == nil {
		sp.categoryRepo = categoryPGRepository.NewRepository(db)
	}

	return sp.categoryRepo
}

func (sp *serviceProvider) SubcategoryRepository(_ context.Context) repository.SubcategoryRepository {
	// TODO: ПЕРЕНЕСТИ!
	db, err := sql.Open("postgres", sp.PGConfig().DSN())
	if err != nil {
		log.Fatal(err)
	}
	// TODO: defer db.Close()

	// TODO: ПЕРЕНЕСТИ!
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected!")

	if sp.subcategoryRepo == nil {
		sp.subcategoryRepo = subcategoryPGRepository.NewRepository(db)
	}

	return sp.subcategoryRepo
}

func (sp *serviceProvider) QuestionService(ctx context.Context) service.QuestionService {
	if sp.questionServ == nil {
		sp.questionServ = questionService.NewService(sp.QuestionRepository(ctx))
	}

	return sp.questionServ
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
