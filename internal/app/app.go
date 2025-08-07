package app

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/Kosfedev/learn_go/internal/config"
	categoryDesc "github.com/Kosfedev/learn_go/pkg/category_v1"
	domainDesc "github.com/Kosfedev/learn_go/pkg/domain_v1"
	questionDesc "github.com/Kosfedev/learn_go/pkg/question_v1"
	setDesc "github.com/Kosfedev/learn_go/pkg/set_v1"
	subcategoryDesc "github.com/Kosfedev/learn_go/pkg/subcategory_v1"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}

	err := app.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (app *App) initDeps(ctx context.Context) error {
	inits := []func(ctx2 context.Context) error{
		app.initConfig,
		app.initServiceProvider,
		app.initGRPCServer,
	}

	for _, init := range inits {
		if err := init(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (app *App) initConfig(_ context.Context) error {
	// TODO: добавить ввод с консоли
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (app *App) initServiceProvider(_ context.Context) error {
	app.serviceProvider = newServiceProvider()
	return nil
}

func (app *App) initGRPCServer(ctx context.Context) error {
	app.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(app.grpcServer)
	questionDesc.RegisterQuestionV1Server(app.grpcServer, app.serviceProvider.QuestionImplementation(ctx))
	setDesc.RegisterSetV1Server(app.grpcServer, app.serviceProvider.QuestionSetImplementation(ctx))
	domainDesc.RegisterDomainV1Server(app.grpcServer, app.serviceProvider.DomainImplementation(ctx))
	categoryDesc.RegisterCategoryV1Server(app.grpcServer, app.serviceProvider.CategoryImplementation(ctx))
	subcategoryDesc.RegisterSubcategoryV1Server(app.grpcServer, app.serviceProvider.SubcategoryImplementation(ctx))

	return nil
}

func (app *App) RunGRPCServer() error {
	address := app.serviceProvider.GRPCConfig().Address()
	log.Printf("Starting gRPC server at %s", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	err = app.grpcServer.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}
