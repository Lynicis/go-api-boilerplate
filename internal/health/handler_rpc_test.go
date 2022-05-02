package health

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	health_proto "go-rest-api-boilerplate/internal/health/proto/health"
)

func Test_RegisterHealthCheckService(t *testing.T) {
	var err error

	testRPCServer := grpc.NewServer()
	port := ":8080"

	RegisterHealthCheckService(testRPCServer)

	ctx := context.Background()
	connection, err := grpc.Dial(
		port,
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)

	defer func(connection *grpc.ClientConn) {
		err = connection.Close()
		assert.NoError(t, err)
	}(connection)

	client := health_proto.NewHealthCheckServiceClient(connection)

	var actualResponse *health_proto.HealthCheckResponse
	actualResponse, err = client.HealthCheck(ctx, &health_proto.HealthCheckRequest{})

	assert.NoError(t, err)
	assert.Equal(t, "OK", actualResponse.Status)
}

func TestRpcHealthService_HealthCheck(t *testing.T) {
	testRPCServer := grpc.NewServer()
	service := RegisterHealthCheckService(testRPCServer)
	status, _ := service.HealthCheck(context.Background(), &health_proto.HealthCheckRequest{})
	assert.IsType(t, status, &health_proto.HealthCheckResponse{})
}
