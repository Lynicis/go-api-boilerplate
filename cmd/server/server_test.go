//go:build unit

package server

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	configmock "go-rest-api-boilerplate/pkg/config/mock"
)

func TestNewServer(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	testConfig := configmock.NewMockConfig(mockController)

	t.Run("should create server instance and return server instance", func(t *testing.T) {
		testServer := NewServer(testConfig)

		expected := &server{}

		assert.IsType(t, expected, testServer)
	})

	t.Run("should server start and stop without error", func(t *testing.T) {
		testServer := NewServer(testConfig)
		go func() {
			err := testServer.Start()
			assert.Nil(t, err)
		}()

		go func() {
			err := testServer.Stop()
			assert.Nil(t, err)
		}()
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
