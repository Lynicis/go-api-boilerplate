//go:build unit

package rpcserver

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	configmodel "go-rest-api-boilerplate/pkg/config/model"
	health "go-rest-api-boilerplate/pkg/rpc_server/testdata"
)

func TestNewRPCServer(t *testing.T) {
	t.Run("should create new rpcServer server and return new rpcServer server instance", func(t *testing.T) {
		rpcServerConfig := configmodel.RPCServer{
			Port: 8091,
		}

		rpcServerInstance := NewRPCServer(rpcServerConfig)

		expectedRPCServerInstance := &rpcServer{}

		assert.NotNil(t, rpcServerInstance)
		assert.IsType(t, expectedRPCServerInstance, rpcServerInstance)
	})

	t.Run("should start rpc server without error", func(t *testing.T) {
		rpcServerConfig := configmodel.RPCServer{
			Port: 8080,
		}

		testRPCServer := NewRPCServer(rpcServerConfig)
		rpcServerInstance := testRPCServer.GetRPCServer()

		RegisterHealthCheckService(rpcServerInstance)

		go func() {
			err := testRPCServer.Start()
			assert.Nil(t, err)

			testRPCServer.Stop()
		}()

		time.Sleep(3 * time.Second)

		connection, err := grpc.Dial(
			fmt.Sprintf(":%d", rpcServerConfig.Port),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)

		defer func(connection *grpc.ClientConn) {
			err = connection.Close()
			assert.Nil(t, err)
		}(connection)

		client := health.NewHealthCheckServiceClient(connection)
		ctx := context.Background()

		response, err := client.HealthCheck(ctx, &health.HealthCheckRequest{})
		expectedResponse := &health.HealthCheckResponse{
			Status: "OK",
		}

		assert.Nil(t, err)
		assert.IsType(t, expectedResponse.Status, response.Status)
	})

	t.Run("should start rpc server return error", func(t *testing.T) {
		rpcServerConfig := configmodel.RPCServer{
			Port: -1100,
		}

		testRPCServer := NewRPCServer(rpcServerConfig)

		go func() {
			err := testRPCServer.Start()
			assert.NotNil(t, err)

			testRPCServer.Stop()
		}()
	})
}

func TestRpcServer_GetRPCServer(t *testing.T) {
	rpcServerConfig := configmodel.RPCServer{
		Port: 8091,
	}

	newRPCServer := NewRPCServer(rpcServerConfig)

	rpcServerInstance := newRPCServer.GetRPCServer()
	expected := &grpc.Server{}

	assert.IsType(t, expected, rpcServerInstance)
}
