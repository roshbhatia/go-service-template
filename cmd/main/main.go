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

	handlers := &api.Api{
		Logger: *logger,
	}
	http.HandleFunc("/health", handlers.HealthCheckHandler)
	http.HandleFunc("/echo", handlers.EchoHandler)

	logger.Fatal.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil))
}
