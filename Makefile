include .env
LOCAL_BIN:=$(CURDIR)/bin

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

install-grpc:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.5
install-minimock:
	GOBIN=$(LOCAL_BIN) go install github.com/gojuno/minimock/v3/cmd/minimock@v3.4.5
install-goose:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.24.3
install-deps:
	make install-grpc
	make install-golangci-lint
	make install-minimock
	make install-goose

generate-question-api:
	mkdir -p pkg/question_v1
	protoc --proto_path api/question_v1 \
	--go_out=pkg/question_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/question_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/question_v1/question.proto
generate-domain-api:
	mkdir -p pkg/domain_v1
	protoc --proto_path api/domain_v1 \
	--go_out=pkg/domain_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/domain_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/domain_v1/domain.proto
generate-category-api:
	mkdir -p pkg/category_v1
	protoc --proto_path api/category_v1 \
	--go_out=pkg/category_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/category_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/category_v1/category.proto
generate-subcategory-api:
	mkdir -p pkg/subcategory_v1
	protoc --proto_path api/subcategory_v1 \
	--go_out=pkg/subcategory_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/subcategory_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/subcategory_v1/subcategory.proto

generate:
	make generate-question-api
	make generate-domain-api
	make generate-category-api
	make generate-subcategory-api

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

init-folders:
	mkdir "tests"
	mkdir "tests/coverage"

test-service:
	go test ./internal/service/question/... -covermode count -coverpkg=./internal/service/question... -cover -coverprofile=tests/coverage/service.out -count 5
	go test ./internal/service/domain/... -covermode count -coverpkg=./internal/service/domain... -cover -coverprofile=tests/coverage/service.out -count 5
	go test ./internal/service/category/... -covermode count -coverpkg=./internal/service/category... -cover -coverprofile=tests/coverage/service.out -count 5
	go test ./internal/service/subcategory/... -covermode count -coverpkg=./internal/service/subcategory... -cover -coverprofile=tests/coverage/service.out -count 5


.PHONY: test
test:
	go clean -testcache
	make test-service

local-migration-status:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DSN} status -v
local-migration-up:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DSN} up -v
local-migration-down:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${PG_DSN} down -v