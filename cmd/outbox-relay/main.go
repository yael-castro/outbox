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

	// Declaration of dependencies
	var cmd command.Command

	// Injecting dependencies
	c := container.New()

	err := c.Inject(ctx, &cmd)
	if err != nil {
		log.Println(err)
		return
	}

	// Listening for shutdown gracefully
	shutdownCh := make(chan struct{}, 1)

	go func() {
		// Waiting for close gracefully
		<-ctx.Done()

		shutdown(c)

		// Confirm shutdown gracefully
		shutdownCh <- struct{}{}
		close(shutdownCh)
	}()

	// Executing command
	exitCodeCh := make(chan int, 1)

	go func() {
		log.Printf("Message relay version '%s' is running", runtime.GitCommit)
		exitCodeCh <- cmd(ctx)
		close(exitCodeCh)
	}()

	// Waiting for cancellation or exit code
	select {
	case <-ctx.Done():
		stop()
		<-shutdownCh

	case exitCode := <-exitCodeCh:
		stop()
		<-shutdownCh

		os.Exit(exitCode)
	}
}

func shutdown(c container.Container) {
	const gracePeriod = 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), gracePeriod)
	defer cancel()

	err := c.Close(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("DI container gracefully closed")
}
