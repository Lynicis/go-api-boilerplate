PROJECT_NAME = $(notdir $(CURDIR))
APP_ENV = local

install_dependencies:
	go get -u ./...
	go mod tidy

lint:
	golangci-lint run ./...

run_local:
	go build -o $(PROJECT_NAME) cmd/main.go
	APP_ENV=$(APP_ENV) ./$(PROJECT_NAME)

run_unit_tests:
	make generate_mock
	go test -tags=unit ./...

generate_mock:
	mockgen -source=pkg/config/config.go -destination=pkg/config/mock/config.go -package=configmock

generate_proto:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative pkg/rpc_server/proto/health.proto