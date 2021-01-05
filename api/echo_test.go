package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/roshbhatia/echo-service/logger"
	"github.com/stretchr/testify/assert"
)

var echoResponse echo

func TestEchoHandler(t *testing.T) {
	logger := logger.NewLogger(os.Stdout)
	ctx, _ := context.WithCancel(context.Background())

	api := &Api{
		Logger: *logger,
		Ctx:    ctx,
	}

	reqBody := strings.NewReader("{\"EchoStr\":\"hello world\"}")
	req, err := http.NewRequest("POST", "/health", reqBody)
	assert.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	echoHandler := http.HandlerFunc(api.EchoHandler)

	echoHandler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Result().StatusCode)

	err = json.Unmarshal(responseRecorder.Body.Bytes(), &echoResponse)
	assert.NoError(t, err)

	assert.Equal(t, "hello world", echoResponse.EchoStr)
	// TODO: test for case when ctx is cancelled
}
