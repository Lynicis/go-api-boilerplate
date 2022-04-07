//go:build unit

package rpcserver

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	configmock "go-rest-api-boilerplate/pkg/config/mock"
	configmodel "go-rest-api-boilerplate/pkg/config/model"
	health "go-rest-api-boilerplate/pkg/rpc_server/proto"
)

func TestNewRPCServer(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("should create new rpcServer server and return new rpcServer server instance", func(t *testing.T) {
		testConfig := configmock.NewMockConfig(mockController)
		rpcServerInstance := NewRPCServer(testConfig)

		expectedRPCServerInstance := &rpcServer{}

		assert.NotNil(t, rpcServerInstance)
		assert.IsType(t, expectedRPCServerInstance, rpcServerInstance)
	})

	t.Run("should start rpc server without error", func(t *testing.T) {
		testConfig := configmock.NewMockConfig(mockController)

		testConfig.EXPECT().GetRPCConfig().Return(configmodel.RPCServer{Port: 8091}).Times(1)

		testRPCServer := NewRPCServer(testConfig)
		rpcServerInstance := testRPCServer.GetRPCServer()

		RegisterHealthCheckService(rpcServerInstance)

		go func() {
			err := testRPCServer.StartServer()
			if err != nil {
				t.Fail()
			}
		}()

		connection, err := grpc.Dial(":8091", grpc.WithInsecure())
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
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	testConfig := configmock.NewMockConfig(mockController)
	newRPCServer := NewRPCServer(testConfig)

	rpcServerInstance := newRPCServer.GetRPCServer()
	expected := &grpc.Server{}

	assert.IsType(t, expected, rpcServerInstance)
}
