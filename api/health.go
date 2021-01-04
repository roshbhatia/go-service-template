package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type healthCheck struct {
	timeStamp, healthStatus, httpStatus string
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// No additional checks that we can do here, as there are no attached resources
	// However, a healthcheck endpoint is valuable as it can be used to register w/ a layer 7 load balancer
	health := healthCheck{
		timeStamp:    time.Now().String(),
		healthStatus: "healthy",
		httpStatus:   http.StatusText(http.StatusOK),
	}
	responseJson, err := json.Marshal(health)
	fmt.Printf("%x\n", responseJson)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseJson))
}
