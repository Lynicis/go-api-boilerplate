//go:build unit

package rpcserver

import (
	"github.com/golang/mock/gomock"
	configmock "go-rest-api-boilerplate/pkg/config/mock"
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

	t.Run("should start rpc server without error", func(t *testing.T) {
		rpcServerInstance := NewRPCServer(testConfig)

		go func() {
			err := rpcServerInstance.StartServer()
			if err != nil {
				t.Fail()
			}
		}()
	})
}
