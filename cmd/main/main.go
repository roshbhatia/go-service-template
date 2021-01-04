package main

import (
	"fmt"
	"net/http"

	"github.com/roshbhatia/echo-service/api"
	"github.com/roshbhatia/echo-service/logger"
)

var (
	port int = 8080
)

func main() {
	logger := logger.NewLogger()
	logger.Info.Printf("starting echo-service on port %d\n", port)

	health := &api.Health{
		Logger: *logger,
	}
	http.HandleFunc("/health", health.HealthCheckHandler)
	logger.Fatal.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil))
}
