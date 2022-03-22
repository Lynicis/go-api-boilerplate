//go:build unit

package healtcheck

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_HealthCheckHandler(t *testing.T) {
	testServer := fiber.New()
	testServer.Get("/health", GetStatus)

	request := httptest.NewRequest(fiber.MethodGet, "/health", nil)
	response, err := testServer.Test(request, 5)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	body, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, "OK", string(body))
}
