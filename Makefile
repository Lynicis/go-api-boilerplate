PROJECT_NAME = $(notdir $(CURDIR))
APP_ENV = local

install_dependencies:
	go get -u ./...
	go mod tidy

lint:
	golangci-lint run ./...

run_local:
	go build -o $(PROJECT_NAME) cmd/main.go
	APP_ENV=$(APP_ENV) go run $(PROJECT_NAME)

run_unit_tests:
	make generate_mock
	go test -tags=unit ./...

build_docker:
	docker build -t $(PROJECT_NAME) .

generate_coverage:
	go test -coverprofile=coverage.txt -covermode=atomic ./...

html_report:
	go test -tags=unit -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out

generate_mock:
	mockgen -source=pkg/logger/logger.go -destination=pkg/logger/mock/logger.go -package=loggermock

generate_proto:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative pkg/rpc_server/proto/health.proto