package health

import (
	"context"

	"google.golang.org/grpc"

	health_proto "go-rest-api-boilerplate/internal/health/proto/health"
)

type RPCHealthService interface {
	HealthCheck(
		context.Context,
		*health_proto.HealthCheckRequest,
	) (
		*health_proto.HealthCheckResponse,
		error,
	)
}

type rpcHealthService struct {
	*health_proto.UnimplementedHealthCheckServiceServer
}

func RegisterHealthCheckService(server grpc.ServiceRegistrar) RPCHealthService {
	service := &rpcHealthService{}
	health_proto.RegisterHealthCheckServiceServer(
		server,
		service,
	)

	return service
}

func (*rpcHealthService) HealthCheck(
	context.Context,
	*health_proto.HealthCheckRequest,
) (
	*health_proto.HealthCheckResponse,
	error,
) {
	return &health_proto.HealthCheckResponse{
		Status: "OK",
	}, nil
}
