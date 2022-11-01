//go:build unit

package server

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"go-rest-api-boilerplate/internal/health"
	"go-rest-api-boilerplate/pkg/config"
)

func Test_NewServer(t *testing.T) { // todo: need mock
	t.Run("should create server instance and return server instance", func(t *testing.T) {
		serverConfig := config.ReadConfig()

		httpServer := NewServer(config)

		serverInstance := &server{}

		assert.NotNil(t, httpServer)
		assert.IsType(t, serverInstance, httpServer)
	})

	t.Run("should server start and stop without error_producer", func(t *testing.T) {
		serverConfig := config_model.HTTPServer{
			Port: 8090,
		}

		testServer := NewServer(serverConfig)

		go func() {
			var err error

			err = testServer.Start()
			assert.NoError(t, err)

			err = testServer.Shutdown()
			assert.NoError(t, err)
		}()

		time.Sleep(3 * time.Second)

		testFiberInstance := testServer.GetFiberInstance()
		registerEndpointForTest(testFiberInstance)

		request := httptest.NewRequest(fiber.MethodGet, "/health", nil)
		response, err := testFiberInstance.Test(request, -1)

		assert.NoError(t, err)
		assert.Equal(t, response.StatusCode, fiber.StatusOK)
	})

	t.Run("should server start and stop return error_producer", func(t *testing.T) {
		serverConfig := config_model.HTTPServer{
			Port: -1000,
		}

		testServer := NewServer(serverConfig)

		go func() {
			err := testServer.Start()
			assert.Error(t, err)

			_ = testServer.Shutdown()
		}()
	})
}

func TestServer_GetFiberInstance(t *testing.T) {
	serverConfig := config_model.HTTPServer{
		Port: 8090,
	}

	httpServerInstance := NewServer(serverConfig)
	getFiberInstance := httpServerInstance.GetFiberInstance()

	fiberInstance := &fiber.App{}

	assert.NotNil(t, httpServerInstance)
	assert.IsType(t, fiberInstance, getFiberInstance)
}

func registerEndpointForTest(fiberInstance *fiber.App) {
	fiberInstance.Get("/health", health.GetStatus)
}
