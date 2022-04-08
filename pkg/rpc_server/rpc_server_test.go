//go:build unit

package rpcserver

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	loggermock "go-rest-api-boilerplate/pkg/logger/mock"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	configmodel "go-rest-api-boilerplate/pkg/config/model"
	health "go-rest-api-boilerplate/pkg/rpc_server/proto"
)

func TestNewRPCServer(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("should create new rpcServer server and return new rpcServer server instance", func(t *testing.T) {
		mockedLogger := loggermock.NewMockLogger(mockController)
		rpcServerConfig := configmodel.RPCServer{
			Port: 8091,
		}

		rpcServerInstance := NewRPCServer(rpcServerConfig, mockedLogger)

		expectedRPCServerInstance := &rpcServer{}

		assert.NotNil(t, rpcServerInstance)
		assert.IsType(t, expectedRPCServerInstance, rpcServerInstance)
	})

	t.Run("should start rpc server without error", func(t *testing.T) {
		mockedLogger := loggermock.NewMockLogger(mockController)
		rpcServerConfig := configmodel.RPCServer{
			Port: 8080,
		}

		testRPCServer := NewRPCServer(rpcServerConfig, mockedLogger)
		rpcServerInstance := testRPCServer.GetRPCServer()

		RegisterHealthCheckService(rpcServerInstance)

		go testRPCServer.Start()

		time.Sleep(5 * time.Second)

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
}

func TestRpcServer_GetRPCServer(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockedLogger := loggermock.NewMockLogger(mockController)
	rpcServerConfig := configmodel.RPCServer{
		Port: 8091,
	}

	newRPCServer := NewRPCServer(rpcServerConfig, mockedLogger)

	rpcServerInstance := newRPCServer.GetRPCServer()
	expected := &grpc.Server{}

	assert.IsType(t, expected, rpcServerInstance)
}
