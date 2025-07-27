package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	categoryImplementation "github.com/Kosfedev/learn_go/internal/api/category"
	domainImplementation "github.com/Kosfedev/learn_go/internal/api/domain"
	quesionImplementation "github.com/Kosfedev/learn_go/internal/api/question"
	subcategoryImplementation "github.com/Kosfedev/learn_go/internal/api/subcategory"
	categoryService "github.com/Kosfedev/learn_go/internal/service/category"
	domainService "github.com/Kosfedev/learn_go/internal/service/domain"
	questionService "github.com/Kosfedev/learn_go/internal/service/question"
	subcategoryService "github.com/Kosfedev/learn_go/internal/service/subcategory"
	categoryDesc "github.com/Kosfedev/learn_go/pkg/category_v1"
	domainDesc "github.com/Kosfedev/learn_go/pkg/domain_v1"
	questionDesc "github.com/Kosfedev/learn_go/pkg/question_v1"
	subcategoryDesc "github.com/Kosfedev/learn_go/pkg/subcategory_v1"
)

const (
	grpcPort = "50051"
)

func main() {
	questionServ := questionService.NewService()
	questionImpl := quesionImplementation.NewImplementation(questionServ)
	domainServ := domainService.NewService()
	domainImpl := domainImplementation.NewImplementation(domainServ)
	categoryServ := categoryService.NewService()
	categoryImpl := categoryImplementation.NewImplementation(categoryServ)
	subcategoryServ := subcategoryService.NewService()
	subcategoryImpl := subcategoryImplementation.NewImplementation(subcategoryServ)
	server := initGRPCServer(questionImpl, domainImpl, categoryImpl, subcategoryImpl)
	address := fmt.Sprintf("localhost:%s", grpcPort)

	log.Printf("Starting gRPC server at %s", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// TODO: change arguments
func initGRPCServer(questionImpl *quesionImplementation.Implementation, domainImpl *domainImplementation.Implementation, categoryImpl *categoryImplementation.Implementation, subcategoryImpl *subcategoryImplementation.Implementation) *grpc.Server {
	server := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(server)
	domainDesc.RegisterDomainV1Server(server, domainImpl)
	questionDesc.RegisterQuestionV1Server(server, questionImpl)
	categoryDesc.RegisterCategoryV1Server(server, categoryImpl)
	subcategoryDesc.RegisterSubcategoryV1Server(server, subcategoryImpl)
	return server
}
