//go:build unit

package health_handler

import (
	"encoding/json"
	health_handler_model "go-rest-api-boilerplate/internal/health_handler/model"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func Test_HealthCheckHandler(t *testing.T) {
	testServer := fiber.New()
	testServer.Get("/health", GetStatus)

	request := httptest.NewRequest(fiber.MethodGet, "/health", nil)
	response, err := testServer.Test(request, 1)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	body, err := ioutil.ReadAll(response.Body)
	assert.Nil(t, err)

	var marshalledBody health_handler_model.HealthEndpoint
	err = json.Unmarshal(body, &marshalledBody)
	assert.Nil(t, err)

	expectedBody := health_handler_model.HealthEndpoint{Status: "OK"}
	assert.Equal(t, expectedBody, marshalledBody)
}
