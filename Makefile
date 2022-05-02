PROJECT_NAME = $(notdir $(CURDIR))
APP_ENV = local

install_dependencies:
	go get ./...
	go mod tidy

lint:
	golangci-lint run ./...

run:
	go run cmd/main.go
	APP_ENV=$(APP_ENV)  ./$(PROJECT_NAME)

test:
	go clean -testcache
	go test -tags=unit ./...

build_docker:
	docker build -t $(PROJECT_NAME) .

coverage_report:
	go clean -testcache
	go test -tags=unit -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out

generate_mock:

generate_proto:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative internal/health/proto/health/health.proto