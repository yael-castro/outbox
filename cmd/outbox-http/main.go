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
	mux := http.NewServeMux()

	if err := container.Inject(ctx, mux); err != nil {
		log.Println(err)
		return
	}

	// Building http server
	server := http.Server{
		Handler: mux,
		Addr:    ":" + port,
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
	}

	// Listening for shutdown gracefully
	doneCh := make(chan struct{})
	defer close(doneCh)

	go shutdown(ctx, &server, doneCh)

	log.Printf("Server http version '%s' is running on port '%s'\n", runtime.GitCommit, port)
	log.Println(server.ListenAndServe())
	<-doneCh
}

func shutdown(ctx context.Context, server *http.Server, doneCh chan<- struct{}) {
	<-ctx.Done()

	defer func() {
		doneCh <- struct{}{}
	}()

	{
		const gracePeriod = 10 * time.Second

		ctx, cancel := context.WithTimeout(context.Background(), gracePeriod)
		defer cancel()

		err := server.Shutdown(ctx)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("Server shutdown gracefully")

		db, err := container.SingleDB(ctx)
		if err != nil {
			log.Println(err)
			return
		}

		err = db.Close()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("Database close gracefully")
	}
}
