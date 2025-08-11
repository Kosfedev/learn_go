package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/Kosfedev/learn_go/internal/app/serviceprovider"
	"github.com/Kosfedev/learn_go/internal/config"
	categoryDesc "github.com/Kosfedev/learn_go/pkg/category_v1"
	domainDesc "github.com/Kosfedev/learn_go/pkg/domain_v1"
	questionFormUpdaterDesc "github.com/Kosfedev/learn_go/pkg/question_form_updater_v1"
	questionFormDesc "github.com/Kosfedev/learn_go/pkg/question_form_v1"
	questionDesc "github.com/Kosfedev/learn_go/pkg/question_v1"
	setDesc "github.com/Kosfedev/learn_go/pkg/set_v1"
	subcategoryDesc "github.com/Kosfedev/learn_go/pkg/subcategory_v1"
)

const (
	swaggerFilePath = "./api/docs/service-api.swagger.json"
	swaggerFileURL  = "/api/docs/service-api.swagger.json"
	swaggerUIURL    = "/api/swagger/"
)

type App struct {
	serviceProvider *serviceprovider.ServiceProvider
	grpcServer      *grpc.Server
	grpcGateway     *http.Server
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
		app.initGRPCGWServer,
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
	app.serviceProvider = serviceprovider.NewServiceProvider()
	return nil
}

func (app *App) initGRPCGWServer(ctx context.Context) error {
	mainMux := http.NewServeMux()

	gwMux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	handlers := []func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error{
		domainDesc.RegisterDomainV1HandlerFromEndpoint,
		categoryDesc.RegisterCategoryV1HandlerFromEndpoint,
	}

	for _, handler := range handlers {
		if err := handler(ctx, gwMux, app.serviceProvider.GRPCConfig().Address(), opts); err != nil {
			log.Fatal(err)
		}
	}

	mainMux.HandleFunc(swaggerFileURL, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, swaggerFilePath)
	})

	mainMux.Handle(swaggerUIURL, httpSwagger.Handler(
		httpSwagger.URL(swaggerFileURL),
	))

	mainMux.Handle("/api/", http.StripPrefix("/api", gwMux))
	mainMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, fmt.Sprintf("%sindex.html", swaggerUIURL), http.StatusFound)
			return
		}
		http.NotFound(w, r)
	})

	app.grpcGateway = &http.Server{
		Addr:    app.serviceProvider.GRPCGWConfig().Address(),
		Handler: mainMux,
	}

	return nil
}

func (app *App) initGRPCServer(ctx context.Context) error {
	app.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(app.grpcServer)
	questionDesc.RegisterQuestionV1Server(app.grpcServer, app.serviceProvider.QuestionImplementation(ctx))
	questionFormDesc.RegisterQuestionFormV1Server(app.grpcServer, app.serviceProvider.QuestionFormImplementation(ctx))
	questionFormUpdaterDesc.RegisterQuestionFormUpdaterV1Server(app.grpcServer, app.serviceProvider.QuestionFormUpdaterImplementation(ctx))
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

func (app *App) RunGRPCGWServer() error {
	log.Printf("Starting HTTP gRPC-Gateway on %s", app.serviceProvider.GRPCGWConfig().Address())
	err := app.grpcGateway.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
