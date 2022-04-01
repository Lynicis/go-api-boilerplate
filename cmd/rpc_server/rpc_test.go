//go:build unit

package rpcserver

import (
	"github.com/golang/mock/gomock"
	configmock "go-rest-api-boilerplate/pkg/config/mock"
	configmodel "go-rest-api-boilerplate/pkg/config/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RPC(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	testConfig := configmock.NewMockConfig(mockController)

	t.Run("should create new rpcServer server and return new rpcServer server instance", func(t *testing.T) {
		rpcServerInstance := NewRPCServer(testConfig)

		expectedRPCServerInstance := &rpcServer{}

		assert.IsType(t, expectedRPCServerInstance, rpcServerInstance)
	})

	t.Run("should start and stop rpcServer server", func(t *testing.T) {
		testConfig.
			EXPECT().
			GetRPCConfig().
			Return(configmodel.RPCServer{Port: 12345}).
			Times(1)

		rpcServerInstance := NewRPCServer(testConfig)

		err := rpcServerInstance.StartServer()
		assert.Nil(t, err)

		rpcServerInstance.GetRPCServer().GracefulStop()
	})
}
