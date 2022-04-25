//go:build unit

package server

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"go-rest-api-boilerplate/internal/healthhandler"
	configmodel "go-rest-api-boilerplate/pkg/config/model"
)

func Test_NewHTTPServer(t *testing.T) {
	t.Run("should create server instance and return server instance", func(t *testing.T) {
		serverConfig := configmodel.HTTPServer{
			Port: 8090,
		}

		httpServerInstance := NewHTTPServer(serverConfig)

		serverInstance := &server{}

		assert.NotNil(t, httpServerInstance)
		assert.IsType(t, serverInstance, httpServerInstance)
	})

	t.Run("should server start and stop without error", func(t *testing.T) {
		serverConfig := configmodel.HTTPServer{
			Port: 8090,
		}

		testServer := NewHTTPServer(serverConfig)

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

	t.Run("should server start and stop return error", func(t *testing.T) {
		serverConfig := configmodel.HTTPServer{
			Port: -1000,
		}

		testServer := NewHTTPServer(serverConfig)

		go func() {
			err := testServer.Start()
			assert.Error(t, err)

			_ = testServer.Shutdown()
		}()
	})
}

func TestServer_GetFiberInstance(t *testing.T) {
	serverConfig := configmodel.HTTPServer{
		Port: 8090,
	}

	httpServerInstance := NewHTTPServer(serverConfig)
	getFiberInstance := httpServerInstance.GetFiberInstance()

	fiberInstance := &fiber.App{}

	assert.NotNil(t, httpServerInstance)
	assert.IsType(t, fiberInstance, getFiberInstance)
}

func registerEndpointForTest(fiberInstance *fiber.App) {
	fiberInstance.Get("/health", healthhandler.GetStatus)
}
