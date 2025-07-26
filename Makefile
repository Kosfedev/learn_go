LOCAL_BIN:=$(CURDIR)/bin

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.5

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

init-folders:
	mkdir "tests/coverage"

test-service:
	go test ./internal/service/... -cover -coverprofile=tests/coverage/service.out

test-service-detailed:
	go test ./internal/service/... -v -cover -coverprofile=tests/coverage/service.out