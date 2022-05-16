//go:build unit

package example_rpc

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	health_proto "go-rest-api-boilerplate/internal/example-rpc/proto"
	"go-rest-api-boilerplate/pkg/config"
	configmodel "go-rest-api-boilerplate/pkg/config/model"
	"go-rest-api-boilerplate/pkg/rpc_server"
)

func TestRegisterHandler(t *testing.T) {
	testRPCServer := grpc.NewServer()
	RegisterHandler(testRPCServer)
	service := &rpcHealthService{}

	status, err := service.HealthCheck(context.Background(), &health_proto.HealthCheckRequest{})

	assert.NoError(t, err)
	assert.IsType(t, status, &health_proto.HealthCheckResponse{})
}

func TestRegisterHandler_HealthCheck(t *testing.T) {
	mainConfig := config.Init(
		configmodel.Fields{
			Server: configmodel.Server{
				RPC: configmodel.RPCServer{
					Port: 8070,
				},
			},
		},
		"local",
	)
	rpcConfig := mainConfig.GetServerConfig().RPC
	rpcServer := rpc_server.NewRPCServer(rpcConfig)
	grpcServer := rpcServer.GetRPCServer()

	go func() {
		err := rpcServer.Start()
		require.NoError(t, err)
	}()

	RegisterHandler(grpcServer)

	time.Sleep(3 * time.Second)

	connection, err := grpc.Dial(
		fmt.Sprintf("127.0.0.1:%d", rpcConfig.Port),
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	require.NoError(t, err)

	defer func(connection *grpc.ClientConn) {
		err = connection.Close()
		require.NoError(t, err)
	}(connection)

	client := health_proto.NewHealthCheckServiceClient(connection)

	ctx := context.Background()
	actualResponse, err := client.HealthCheck(ctx, &health_proto.HealthCheckRequest{})

	assert.NoError(t, err)
	assert.Equal(t, "OK", actualResponse.Status)
}
