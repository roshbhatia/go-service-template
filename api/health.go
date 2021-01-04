package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/roshbhatia/echo-service/logger"
)

type Health struct {
	Logger logger.Logger
}

type healthCheck struct {
	timeStamp, healthStatus, httpStatus string
}

func (h *Health) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// No additional checks that we can do here, as there are no attached resources
	// However, a healthcheck endpoint is valuable as it can be used to register w/ a layer 7 load balancer
	healthResp := healthCheck{
		timeStamp:    time.Now().String(),
		healthStatus: "healthy",
		httpStatus:   http.StatusText(http.StatusOK),
	}
	h.Logger.Info.Printf("client %s requested health status", r.RemoteAddr)
	responseJson, err := json.Marshal(healthResp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.Logger.Err.Printf("failed to marshal heath check object")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseJson))
}
