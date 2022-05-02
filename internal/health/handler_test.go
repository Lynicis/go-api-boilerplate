//go:build unit

package health

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	health_handler_model "go-rest-api-boilerplate/internal/health/model"
)

func Test_HealthCheckHandler(t *testing.T) {
	testHTTPServer := fiber.New()
	testHTTPServer.Get("/health", GetStatus)

	request := httptest.NewRequest(fiber.MethodGet, "/health", nil)
	response, err := testHTTPServer.Test(request, 1)

	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	body, err := ioutil.ReadAll(response.Body)

	var marshalledBody health_handler_model.HealthEndpoint
	err = json.Unmarshal(body, &marshalledBody)

	expectedBody := health_handler_model.HealthEndpoint{
		Status: "OK",
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedBody, marshalledBody)
}
