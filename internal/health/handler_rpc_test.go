//go:build unit

package health

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"

	health_proto "go-rest-api-boilerplate/internal/health/proto/health"
)

func TestRpcHealthService(t *testing.T) {
	testRPCServer := grpc.NewServer()
	service := RegisterHealthCheckService(testRPCServer)
	status, _ := service.HealthCheck(context.Background(), &health_proto.HealthCheckRequest{})
	assert.IsType(t, status, &health_proto.HealthCheckResponse{})
}

func Test_RegisterHealthCheckService_HealthCheck(t *testing.T) {
	testRPCServer := grpc.NewServer()
	RegisterHealthCheckService(testRPCServer)

	ctx := context.Background()
	connection, err := grpc.Dial(
		":8080",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	defer func(connection *grpc.ClientConn) {
		err = connection.Close()
		require.NoError(t, err)
	}(connection)

	client := health_proto.NewHealthCheckServiceClient(connection)

	var actualResponse *health_proto.HealthCheckResponse
	actualResponse, err = client.HealthCheck(ctx, &health_proto.HealthCheckRequest{})

	assert.NoError(t, err)
	assert.Equal(t, "OK", actualResponse.Status)
}
