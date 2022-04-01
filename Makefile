install_dependencies:
	go get ./...
	go mo tidy

run_unit_tests:
	go test -tags=unit ./...

generate_mock:
	mockgen -source=pkg/config/config.go -destination=pkg/config/mock/config.go -package=configmock