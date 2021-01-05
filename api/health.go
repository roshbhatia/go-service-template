package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type healthCheck struct {
	TimeStamp, HealthStatus, HttpStatus string
}

func (h *Api) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	threadWorkDone := make(chan bool)

	// No additional checks that we can do here, as there are no attached resources
	// However, a healthcheck endpoint is valuable as it can be used to register w/ a layer 7 load balancer
	go func() {
		healthResp := healthCheck{
			TimeStamp:    time.Now().UTC().String(),
			HealthStatus: "healthy",
			HttpStatus:   http.StatusText(http.StatusOK),
		}
		h.Logger.Info(fmt.Sprintf("client %s requested health status", r.RemoteAddr))

		responseBody, err := json.Marshal(healthResp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			h.Logger.Err("failed to marshal health check object")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(responseBody)
		threadWorkDone <- true
	}()

	// Not really safe per-say, but if the contex is done, we'll just abort the above thread
	select {
	case <-h.Ctx.Done():
		return
	case <-threadWorkDone:
		return
	}
}
