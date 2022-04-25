package rpcserver

import (
	"context"

	"google.golang.org/grpc"

	health "go-rest-api-boilerplate/pkg/rpc_server/testdata"
)

type healthService struct {
	*health.UnimplementedHealthCheckServiceServer
}

func RegisterHealthCheckService(server grpc.ServiceRegistrar) {
	service := &healthService{}

	health.RegisterHealthCheckServiceServer(
		server,
		service,
	)
}

func (*healthService) HealthCheck(
	context.Context,
	*health.HealthCheckRequest,
) (
	*health.HealthCheckResponse,
	error,
) {
	return &health.HealthCheckResponse{
		Status: "OK",
	}, nil
}
