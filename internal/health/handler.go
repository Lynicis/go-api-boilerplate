package health

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"

	health_handler_model "go-rest-api-boilerplate/internal/health/model"
	"go-rest-api-boilerplate/internal/health/proto/health"
)

type rpcHealthService struct {
	*health.UnimplementedHealthCheckServiceServer
}

func GetStatus(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(
		health_handler_model.HealthEndpoint{
			Status: "OK",
		},
	)
}

func RegisterHealthCheckService(server grpc.ServiceRegistrar) {
	service := &rpcHealthService{}

	health.RegisterHealthCheckServiceServer(
		server,
		service,
	)
}

func (*rpcHealthService) HealthCheck(
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
