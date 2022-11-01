//go:build unit

package health

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func Test_HealthCheckHandler(t *testing.T) {
	testHTTPServer := fiber.New()
	testHTTPServer.Get("/health", GetStatus)

	request := httptest.NewRequest(fiber.MethodGet, "/health", nil)
	response, err := testHTTPServer.Test(request, 1)

	assert.Equal(t, fiber.StatusOK, response.StatusCode)
	assert.NoError(t, err)
}
