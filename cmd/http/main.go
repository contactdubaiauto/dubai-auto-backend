package main

import (
	"context"
	app "dubai-auto/internal"
	"dubai-auto/internal/config"
	"dubai-auto/internal/storage/postgres"
	"dubai-auto/pkg"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// TODO: all wrong info send 400 status, ingo log and error log must be seperately
	conf := config.Init()
	pkg.Init(conf.ACCESS_KEY, conf.ACCESS_TIME, conf.REFRESH_KEY, conf.REFRESH_TIME)
	logger := config.InitLogger(conf.LOGGER_FOLDER_PATH, conf.LOGGER_FILENAME, conf.GIN_MODE)
	db := postgres.Init()
	server := app.InitApp(db, conf)
	// utils.InitCron(logger)

	srv := &http.Server{
		Addr:    conf.PORT,
		Handler: server,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
		}
	}()

	// When a shutdown signal (like Ctrl+C) is caught, you initiate a graceful shutdown using srv.Shutdown(ctx), giving active connections time to complete before the server shuts down.
	// Graceful shutdown is handled within the main function, allowing ongoing requests to finish processing before the server shuts down, if req not completed in 5 seconds the server will force shutdown.
	// If the expected request finishes before 5 seconds, the server is shut down immediately.
	// New Requests will not be accepted.
	// Wait for an interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Println("Shutting down server...")

	// Create a context with a timeout for the graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt to gracefully shutdown the server
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown:", err)
	}

	logger.Println("Server exiting")
}
