package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	quesionImplementation "github.com/Kosfedev/learn_go/internal/api/question"
	questionService "github.com/Kosfedev/learn_go/internal/service/question"
	desc "github.com/Kosfedev/learn_go/pkg/question_v1"
)

const (
	grpcPort = "50051"
)

func main() {
	questionServ := questionService.NewService()
	questionImpl := quesionImplementation.NewImplementation(questionServ)
	server := initGRPCServer(questionImpl)
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

func initGRPCServer(questionImpl *quesionImplementation.Implementation) *grpc.Server {
	server := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(server)
	desc.RegisterQuestionV1Server(server, questionImpl)
	return server
}
