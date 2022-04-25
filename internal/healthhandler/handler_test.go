//go:build unit

package healthhandler

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	healthhandlermodel "go-rest-api-boilerplate/internal/healthhandler/model"
)

func Test_HealthCheckHandler(t *testing.T) {
	testHTTPServer := fiber.New()
	testHTTPServer.Get("/health", GetStatus)

	request := httptest.NewRequest(fiber.MethodGet, "/health", nil)
	response, err := testHTTPServer.Test(request, 1)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, response.StatusCode)

	body, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)

	var marshalledBody healthhandlermodel.HealthEndpoint
	err = json.Unmarshal(body, &marshalledBody)
	assert.NoError(t, err)

	expectedBody := healthhandlermodel.HealthEndpoint{
		Status: "OK",
	}

	assert.Equal(t, expectedBody, marshalledBody)
}
