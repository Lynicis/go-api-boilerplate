//go:build unit

package server

import (
	"fmt"
	"github.com/golang/mock/gomock"
	configmock "go-rest-api-boilerplate/pkg/config/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Server(t *testing.T) {
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
			if err != nil {
				fmt.Println(err)
				t.Fail()
			}
		}()
	})
}
