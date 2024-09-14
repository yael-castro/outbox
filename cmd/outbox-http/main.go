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
	port := os.Getenv("PORT")
	if len(port) == 0 {
		const defaultPort = "8080"
		port = defaultPort
	}

	// Building main context
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
	defer stop()

	// Injecting dependencies
	var handler http.Handler

	c := container.New()

	if err := c.Inject(ctx, &handler); err != nil {
		log.Println(err)
		return
	}

	// Building http server
	server := http.Server{
		Handler: handler,
		Addr:    ":" + port,
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
	}

	doneCh := make(chan struct{})
	defer close(doneCh)

	// Listening for shutdown gracefully
	go func() {
		<-ctx.Done()
		shutdown(&server, c, doneCh)
	}()

	log.Printf("Server http version '%s' is running on port '%s'\n", runtime.GitCommit, port)
	log.Println(server.ListenAndServe())

	// Gracefully shutdown
	stop() // TODO: avoid repeat this call
	<-doneCh
}

func shutdown(server *http.Server, c container.Container, doneCh chan<- struct{}) {
	defer func() {
		doneCh <- struct{}{}
	}()

	const gracePeriod = 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), gracePeriod)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Server shutdown gracefully")

	err = c.Close(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Database close gracefully")
}
