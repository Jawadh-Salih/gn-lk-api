package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Jawadh-Salih/gn-lk-api/internal/api"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router
	router := api.InitRoutes()

	// Define a simple handler
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// Create an HTTP server
	srv := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	// Channel to listen for interrupt signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTSTP)

	// Goroutine to handle the signal
	go func() {
		sig := <-sigCh
		fmt.Printf("Received signal: %v\n", sig)

		// Create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Attempt to gracefully shutdown the server.
		if err := srv.Shutdown(ctx); err != nil {
			fmt.Println("Server forced to shutdown:", err)
		} else {
			fmt.Println("Server gracefully stopped")
		}
	}()

	// Start the server
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Listen and serve error: %v\n", err)
	}
}
