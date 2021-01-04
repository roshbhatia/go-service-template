package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type healthCheck struct {
	TimeStamp, HealthStatus, HttpStatus string
}

func (h *Api) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// No additional checks that we can do here, as there are no attached resources
	// However, a healthcheck endpoint is valuable as it can be used to register w/ a layer 7 load balancer
	healthResp := healthCheck{
		TimeStamp:    time.Now().UTC().String(),
		HealthStatus: "healthy",
		HttpStatus:   http.StatusText(http.StatusOK),
	}
	h.Logger.Info.Printf("client %s requested health status", r.RemoteAddr)

	responseBody, err := json.Marshal(healthResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.Logger.Err.Println("failed to marshal health check object")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}
