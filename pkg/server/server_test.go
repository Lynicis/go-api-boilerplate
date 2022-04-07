//go:build unit

package server

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	configmodel "go-rest-api-boilerplate/pkg/config/model"
)

func TestNewServer(t *testing.T) {
	t.Run("should create server instance and return server instance", func(t *testing.T) {
		serverConfig := configmodel.Server{
			Port: 8090,
		}

		testServer := NewServer(serverConfig)

		expected := &server{}

		assert.NotNil(t, testServer)
		assert.IsType(t, expected, testServer)
	})

	t.Run("should server start without error", func(t *testing.T) {
		serverConfig := configmodel.Server{
			Port: 8090,
		}

		testServer := NewServer(serverConfig)

		go func() {
			err := testServer.Start()
			if err != nil {
				t.Fail()
			}
		}()

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
	serverConfig := configmodel.Server{
		Port: 8090,
	}

	testServer := NewServer(serverConfig)
	fiberInstance := testServer.GetFiberInstance()

	expected := &fiber.App{}

	assert.IsType(t, expected, fiberInstance)
}
