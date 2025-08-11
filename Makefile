include .env
LOCAL_BIN:=$(CURDIR)/bin

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

# TODO: добавить клонирование google api && protoc-gen-openapiv2

install-grpc:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.27.1
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.27.1

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

# TODO: перейти на buf и избавиться от этого всего
generate-question-api:
	mkdir -p pkg/question_v1
	protoc --proto_path api/question_v1 --proto_path third_party/googleapis --proto_path third_party/grpc-gateway \
	--go_out=pkg/question_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/question_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/question_v1/question.proto
generate-question-set-api:
	mkdir -p pkg/set_v1
	protoc --proto_path api/set_v1 --proto_path third_party/googleapis --proto_path third_party/grpc-gateway \
	--go_out=pkg/set_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/set_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/set_v1/set.proto
generate-question-form-api:
	mkdir -p pkg/question_form_v1
	protoc --proto_path api/question_form_v1 --proto_path api/question_v1 --proto_path api/set_v1 --proto_path api/subcategory_v1 --proto_path third_party/googleapis --proto_path third_party/grpc-gateway \
	--go_out=pkg/question_form_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/question_form_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/question_form_v1/question_form.proto
generate-question-form-updater-api:
	mkdir -p pkg/question_form_updater_v1
	protoc --proto_path api/question_form_updater_v1 --proto_path api/question_v1 --proto_path third_party/googleapis --proto_path third_party/grpc-gateway \
	--go_out=pkg/question_form_updater_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/question_form_updater_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/question_form_updater_v1/question_form_updater.proto
generate-domain-api:
	mkdir -p pkg/domain_v1
	protoc --proto_path api/domain_v1 --proto_path third_party/googleapis --proto_path third_party/grpc-gateway \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
	--go_out=pkg/domain_v1 --go_opt=paths=source_relative \
	--go-grpc_out=pkg/domain_v1 --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pkg/domain_v1 --grpc-gateway_opt=paths=source_relative --grpc-gateway_opt=generate_unbound_methods=true \
	api/domain_v1/domain.proto
generate-category-api:
	mkdir -p pkg/category_v1
	protoc --proto_path api/category_v1 --proto_path third_party/googleapis --proto_path third_party/grpc-gateway \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
	--go_out=pkg/category_v1 --go_opt=paths=source_relative \
	--go-grpc_out=pkg/category_v1 --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pkg/category_v1 --grpc-gateway_opt=paths=source_relative --grpc-gateway_opt=generate_unbound_methods=true \
	api/category_v1/category.proto
generate-subcategory-api:
	mkdir -p pkg/subcategory_v1
	protoc --proto_path api/subcategory_v1 --proto_path third_party/googleapis --proto_path third_party/grpc-gateway \
	--go_out=pkg/subcategory_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/subcategory_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/subcategory_v1/subcategory.proto
generate-openapi:
	mkdir -p pkg/docs
	protoc \
	--proto_path third_party/googleapis \
	--proto_path third_party/grpc-gateway \
	--proto_path api/domain_v1 \
	--proto_path api/category_v1 \
	--proto_path api/subcategory_v1 \
	--proto_path api/question_v1 \
	--proto_path api/question_form_v1 \
	--proto_path api/question_form_updater_v1 \
	--proto_path api/set_v1 \
	--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
	--openapiv2_out=api/docs --openapiv2_opt=logtostderr=true --openapiv2_opt=allow_merge=true --openapiv2_opt=merge_file_name=service-api \
	api/domain_v1/domain.proto \
	api/category_v1/category.proto \
	api/subcategory_v1/subcategory.proto \
	api/question_v1/question.proto \
	api/question_form_v1/question_form.proto \
	api/question_form_updater_v1/question_form_updater.proto \
	api/set_v1/set.proto

generate:
	make generate-question-api
	make generate-question-set-api
	make generate-question-form-api
	make generate-question-form-updater-api
	make generate-domain-api
	make generate-category-api
	make generate-subcategory-api
	make generate-openapi

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

test-service:
	go test ./internal/service/question/... -coverpkg=./internal/service/question... -count 5
	go test ./internal/service/set/... -coverpkg=./internal/service/set... -count 5
	go test ./internal/service/domain/... -coverpkg=./internal/service/domain... -count 5
	go test ./internal/service/category/... -coverpkg=./internal/service/category... -count 5
	go test ./internal/service/subcategory/... -coverpkg=./internal/service/subcategory... -count 5

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