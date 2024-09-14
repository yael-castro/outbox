package main

import (
	"context"
	"github.com/yael-castro/outbox/internal/app/input/command"
	"github.com/yael-castro/outbox/internal/container"
	"github.com/yael-castro/outbox/internal/runtime"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Building main context
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
	defer stop()

	// DI in action!
	var cmd command.Command

	c := container.New()

	err := c.Inject(ctx, &cmd)
	if err != nil {
		log.Println(err)
		return
	}

	// Listening for shutdown gracefully
	doneCh := make(chan struct{})
	defer close(doneCh)

	go func() {
		<-ctx.Done()
		shutdown(c, doneCh)
	}()

	// Executing command
	log.Printf("Message relay version '%s' is running", runtime.GitCommit)
	cmd(ctx)

	// Gracefully shutdown
	stop() // TODO: avoid repeat this call
	<-doneCh
}

func shutdown(c container.Container, doneCh chan struct{}) {
	defer func() {
		doneCh <- struct{}{}
	}()

	const gracePeriod = 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), gracePeriod)
	defer cancel()

	err := c.Close(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Database close gracefully")
}
