package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 1. Set your routes
	mux := http.NewServeMux()
	mux.HandleFunc("/work", func (w http.ResponseWriter, r *http.Request)  {
		log.Println("Starting work...")
		// Simulate long running task - database query
		time.Sleep(5*time.Second)
		fmt.Fprintln(w, "Work completed successfully")
		log.Println("Finished work.")
	})

	// 2. Configure the http server
	srv := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	// 3. Start the server in separate go routine so it doesn't block
	go func() {
		log.Printf("Starting server on port %s", srv.Addr)
		// ErrServerClosed is expected when we call srv.Shutdown()
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// 4. Set up a channel to listen for OS signal
	quit := make(chan os.Signal, 1)

	// Catch SIGINT (Ctrl+C) and SIGTERM (Docker/Kubernetes standard shutdown)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 5. Block the goroutine until a signal is received.
	<-quit
	log.Println("Shutdown signal recieved. Stopping server...")

	// 6. Create a deadline to wait for in-flight requests to finish.
	// Without this a server could hang forever if a connection is stuck
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 7. Execute the graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown due to error or timeout: %v", err)
	}
	log.Println("Server exited properly.")
}