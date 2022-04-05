//go:build unit

package server

import (
	configmodel "go-rest-api-boilerplate/pkg/config/model"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	configmock "go-rest-api-boilerplate/pkg/config/mock"
)

func TestNewServer(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	t.Run("should create server instance and return server instance", func(t *testing.T) {
		testConfig := configmock.NewMockConfig(mockController)
		testServer := NewServer(testConfig)

		expected := &server{}

		assert.NotNil(t, testServer)
		assert.IsType(t, expected, testServer)
	})

	t.Run("should server start without error", func(t *testing.T) {
		testConfig := configmock.NewMockConfig(mockController)
		testConfig.EXPECT().GetServerConfig().Return(configmodel.Server{Port: 8090})

		testServer := NewServer(testConfig)

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
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	testConfig := configmock.NewMockConfig(mockController)

	testServer := NewServer(testConfig)
	fiberInstance := testServer.GetFiberInstance()

	expected := &fiber.App{}

	assert.IsType(t, expected, fiberInstance)
}
