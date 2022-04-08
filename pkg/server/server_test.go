//go:build unit

package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	loggermock "go-rest-api-boilerplate/pkg/logger/mock"
	"net/http/httptest"
	"testing"

	configmodel "go-rest-api-boilerplate/pkg/config/model"
)

func TestNewServer(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("should create server instance and return server instance", func(t *testing.T) {
		mockedLogger := loggermock.NewMockLogger(mockController)
		serverConfig := configmodel.Server{
			Port: 8090,
		}

		testServer := NewServer(serverConfig, mockedLogger)

		expected := &server{}

		assert.NotNil(t, testServer)
		assert.IsType(t, expected, testServer)
	})

	t.Run("should server start without error", func(t *testing.T) {
		mockedLogger := loggermock.NewMockLogger(mockController)
		serverConfig := configmodel.Server{
			Port: 8090,
		}

		testServer := NewServer(serverConfig, mockedLogger)

		go testServer.Start()

		testFiberInstance := testServer.GetFiberInstance()
		testFiberInstance.Get("/exist", func(ctx *fiber.Ctx) error {
			return ctx.SendStatus(fiber.StatusOK)
		})

		request := httptest.NewRequest(fiber.MethodGet, "/exist", nil)

		response, err := testFiberInstance.Test(request, 1)

		assert.Nil(t, err)
		assert.Equal(t, response.StatusCode, fiber.StatusOK)
	})
}

func TestServer_GetFiberInstance(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockedLogger := loggermock.NewMockLogger(mockController)
	serverConfig := configmodel.Server{
		Port: 8090,
	}

	testServer := NewServer(serverConfig, mockedLogger)
	fiberInstance := testServer.GetFiberInstance()

	expected := &fiber.App{}

	assert.IsType(t, expected, fiberInstance)
}
