package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/roshbhatia/echo-service/api"
	"github.com/roshbhatia/echo-service/logger"
)

func main() {
	logger := logger.NewLogger(os.Stdout)
	ctx, cancel := context.WithCancel(context.Background())

	port, err := strconv.Atoi((os.Getenv("SERVICE_PORT")))
	if err != nil {
		logger.Fatal("SERVICE_PORT is unset or invalid, exiting")
	}

	logger.Info(fmt.Sprintf("starting echo-service on port %d\n", port))

	ssl_cert_path := os.Getenv("SSL_CERT_PATH")
	ssl_key_path := os.Getenv("SSL_KEY_PATH")
	if ssl_cert_path == "" || ssl_key_path == "" {
		logger.Info("SSL_CERT_PATH and SSL_KEY_PATH are not both set, running in HTTP only")
	} else {
		logger.Info("SSL_CERT_PATH and SSL_KEY_PATH provided, running wtih HTTPS")
	}

	// Not best practice to pass around this in structs w/ member functions, dependency injection would be better here
	// TODO inject dependencies rather than pass context and logger around in a struct
	handlers := &api.Api{
		Logger: *logger,
		Ctx:    ctx,
	}
	http.HandleFunc("/health", handlers.HealthCheckHandler)
	http.HandleFunc("/echo", handlers.EchoHandler)

	// Check for SIGTERM, then cleanly shutdown
	osSignal := make(chan os.Signal)
	signal.Notify(osSignal, os.Interrupt, syscall.SIGINT)
	signal.Notify(osSignal, os.Interrupt, syscall.SIGTERM)
	go func() {
		signal := <-osSignal
		var signalString string
		if signal == syscall.SIGINT {
			signalString = "SIGINT"
		} else if signal == syscall.SIGTERM {
			signalString = "SIGTERM"
		}
		logger.Info(fmt.Sprintf("processs recieved %s, shutting down\n", signalString))
		cancel()
		os.Exit(1)
	}()

	logger.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil).Error())

	// if ssl_cert_path == "" || ssl_key_path == "" {
	// 	logger.Fatal.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
	// } else {
	// 	logger.Fatal.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%d", port), ssl_cert_path, ssl_key_path, nil))
	// }
}
