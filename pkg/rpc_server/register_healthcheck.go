package rpcserver

import (
	"context"

	"google.golang.org/grpc"

	health "go-rest-api-boilerplate/pkg/rpc_server/proto"
)

type healthService struct {
	*health.UnimplementedHealthCheckServiceServer
}

// RegisterHealthCheckService Register gRPC health service
func RegisterHealthCheckService(server grpc.ServiceRegistrar) {
	health.RegisterHealthCheckServiceServer(server, &healthService{})
}

func (h *healthService) HealthCheck(context.Context, *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{Status: "OK"}, nil
}
