PROJECT_NAME = $(notdir $(CURDIR))
APP_ENV = local

install_dependencies:
	go get -u ./...
	go mod tidy

lint:
	golangci-lint run ./...

run:
	go build -o $(PROJECT_NAME) cmd/main.go
	APP_ENV=$(APP_ENV) go run $(PROJECT_NAME)

security:
	gosec -no-fail -fmt html -out results.html ./...
	open results.html

run_unit_tests:
	make generate_mock
	go test -tags=unit ./...

build_docker:
	docker build -t $(PROJECT_NAME) .

coverage:
	go test -coverprofile=coverage.txt -covermode=atomic ./...

coverage_report:
	go test -tags=unit -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out

generate_mock:

generate_proto:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative pkg/rpc_server/proto/health.proto