//go:build http

package main

import (
	"context"
	"github.com/yael-castro/outbox/internal/container"
	"github.com/yael-castro/outbox/internal/runtime"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Building main context
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
	defer stop()

	// Declaration of dependencies
	var handler http.Handler

	// Injecting dependencies
	c := container.New()

	if err := c.Inject(ctx, &handler); err != nil {
		log.Println(err)
		return
	}

	// Getting http port
	port := os.Getenv("PORT")
	if len(port) == 0 {
		const defaultPort = "8080"
		port = defaultPort
	}

	// Building http server
	server := http.Server{
		Handler: handler,
		Addr:    ":" + port,
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
	}

	// Listening for shutdown gracefully
	shutdownCh := make(chan struct{}, 1)

	go func() {
		// Waiting for shutdown gracefully
		<-ctx.Done()

		shutdown(c, &server)

		// Confirm shutdown gracefully
		shutdownCh <- struct{}{}
		close(shutdownCh)
	}()

	// Running http server
	errCh := make(chan error, 1)

	go func() {
		log.Printf("Server http version '%s' is running on port '%s'\n", runtime.GitCommit, port)
		errCh <- server.ListenAndServe()
		close(errCh)
	}()

	// Waiting for cancellation or error
	select {
	case <-ctx.Done():
		stop()
		<-shutdownCh

	case err := <-errCh:
		stop()
		<-shutdownCh

		log.Fatal(err)
	}
}

func shutdown(c container.Container, server *http.Server) {
	const gracePeriod = 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), gracePeriod)
	defer cancel()

	// Closing http server
	err := server.Shutdown(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Server shutdown gracefully")

	// Closing DI container
	err = c.Close(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("DI container gracefully closed")
}
