//go:build unit

package rpc_server

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"go-rest-api-boilerplate/internal/health"
	health_proto "go-rest-api-boilerplate/internal/health/proto/health"
	config_model "go-rest-api-boilerplate/pkg/config/model"
)

func TestNewRPCServer(t *testing.T) {
	t.Run("should create new rpcServer server and return new rpcServer server instance", func(t *testing.T) {
		rpcServerConfig := config_model.RPCServer{
			Port: 8091,
		}

		rpcServerInstance := NewRPCServer(rpcServerConfig)

		expectedRPCServerInstance := &rpcServer{}

		assert.NotNil(t, rpcServerInstance)
		assert.IsType(t, expectedRPCServerInstance, rpcServerInstance)
	})

	t.Run("should start rpc server without error", func(t *testing.T) {
		rpcServerConfig := config_model.RPCServer{
			Port: 8080,
		}

		testRPCServer := NewRPCServer(rpcServerConfig)
		rpcServerInstance := testRPCServer.GetRPCServer()

		health.RegisterHealthCheckService(rpcServerInstance)

		go func() {
			err := testRPCServer.Start()
			assert.Nil(t, err)

			testRPCServer.Stop()
		}()

		time.Sleep(3 * time.Second)

		connection, err := grpc.Dial(
			fmt.Sprintf(":%d", rpcServerConfig.Port),
			grpc.WithTransportCredentials(
				insecure.NewCredentials(),
			),
		)

		defer func(connection *grpc.ClientConn) {
			err = connection.Close()
			assert.NoError(t, err)
		}(connection)

		client := health_proto.NewHealthCheckServiceClient(connection)
		ctx := context.Background()

		response, err := client.HealthCheck(ctx, &health_proto.HealthCheckRequest{})
		expectedResponse := &health_proto.HealthCheckResponse{
			Status: "OK",
		}

		assert.NoError(t, err)
		assert.IsType(t, expectedResponse.Status, response.Status)
	})

	t.Run("should start rpc server return error", func(t *testing.T) {
		rpcServerConfig := config_model.RPCServer{
			Port: -1100,
		}

		testRPCServer := NewRPCServer(rpcServerConfig)

		go func() {
			err := testRPCServer.Start()
			assert.Error(t, err)

			testRPCServer.Stop()
		}()
	})
}

func TestRPCServer_GetRPCServer(t *testing.T) {
	rpcServerConfig := config_model.RPCServer{
		Port: 8091,
	}

	newRPCServer := NewRPCServer(rpcServerConfig)

	rpcServerInstance := newRPCServer.GetRPCServer()
	expected := &grpc.Server{}

	assert.IsType(t, expected, rpcServerInstance)
}
