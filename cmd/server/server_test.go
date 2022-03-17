package server

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Server(t *testing.T) {
	t.Run("should create server and return server", func(t *testing.T) {
		testPort := ":8080"
		testServer := NewGatewayServer(testPort)

		expected := &server{
			port:  testPort,
			fiber: testServer.GetFiberInstance(),
		}

		assert.Equal(t, expected, testServer)
	})

	t.Run("should server start and stop without error", func(t *testing.T) {
		testPort := ":8080"
		testServer := NewGatewayServer(testPort)

		go func() {
			err := testServer.Start()
			if err != nil {
				fmt.Println(err)
				t.Fail()
			}
		}()
	})
}
