//go:build unit

package rpcserver

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	configmodel "go-rest-api-boilerplate/pkg/config/model"
	health "go-rest-api-boilerplate/pkg/rpc_server/proto"
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
			Port: 8091,
		}

		testRPCServer := NewRPCServer(rpcServerConfig)
		rpcServerInstance := testRPCServer.GetRPCServer()

		RegisterHealthCheckService(rpcServerInstance)

		go func() {
			err := testRPCServer.StartServer()
			if err != nil {
				t.Fail()
			}
		}()

		connection, err := grpc.Dial(fmt.Sprintf(":%d", rpcServerConfig.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
		defer func(connection *grpc.ClientConn) {
			err = connection.Close()
			assert.Nil(t, err)
		}(connection)

		client := health.NewHealthCheckServiceClient(connection)
		ctx := context.Background()

		response, err := client.HealthCheck(ctx, &health.HealthCheckRequest{})

		assert.Nil(t, err)
		assert.NotNil(t, response.Status, "OK")
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
