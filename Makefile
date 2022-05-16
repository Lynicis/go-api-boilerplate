PROJECT_NAME = $(notdir $(CURDIR))

.PHONY: install_dependencies
get:
	go get ./...
	go mod tidy

.PHONY: lint
lint:
	golangci-lint run -c .golangci.yml ./...

.PHONY: run
run:
	go run cmd/*.go

.PHONY: test
test:
	go clean -testcache
	go test -tags=unit ./...

.PHONY: dockerize
dockerize:
	docker build -t $(PROJECT_NAME) .

.PHONY: coverage_report
coverage_report:
	go clean -testcache
	go test -tags=unit -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out

.PHONY: generate_mock
generate_mock:

.PHONY: generate_proto
generate_proto:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative internal/example-rpc/proto/health.proto