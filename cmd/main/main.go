package main

import (
	"log"
	"net/http"

	"github.com/roshbhatia/echo-service/api"
)

func main() {
	http.HandleFunc("/health", api.HealthCheckHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
