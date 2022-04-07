install_dependencies:
	go get ./...
	go mo tidy

lint:
	golangci-lint run ./...

run_unit_tests:
	make generate_mock
	go test -tags=unit ./...

generate_mock:
	mockgen -source=pkg/config/config.go -destination=pkg/config/mock/config.go -package=configmock

generate_proto:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative pkg/rpc_server/proto/health.proto