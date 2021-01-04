package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/roshbhatia/echo-service/logger"
	"github.com/stretchr/testify/assert"
)

var healthCheckResponse healthCheck

func TestHealthCheckHandler(t *testing.T) {
	logger := logger.NewLogger()
	api := &Api{
		Logger: *logger,
	}

	req, err := http.NewRequest("GET", "/health", nil)
	assert.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	healthCheckHandler := http.HandlerFunc(api.HealthCheckHandler)

	healthCheckHandler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Result().StatusCode)

	err = json.Unmarshal(responseRecorder.Body.Bytes(), &healthCheckResponse)
	assert.NoError(t, err)

	assert.Equal(t, "healthy", healthCheckResponse.HealthStatus)
	assert.Equal(t, "OK", healthCheckResponse.HttpStatus)
}
