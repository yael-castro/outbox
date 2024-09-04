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

	// Listening for shutdown gracefully
	doneCh := make(chan struct{})
	defer close(doneCh)

	go shutdown(ctx, doneCh)

	// DI in action!
	var cmd command.Command

	err := container.Inject(ctx, &cmd)
	if err != nil {
		log.Println(err)
		return
	}

	// Executing command
	log.Printf("Message relay version '%s' is running", runtime.GitCommit)
	cmd(ctx)
	<-doneCh
}

func shutdown(ctx context.Context, doneCh chan struct{}) {
	<-ctx.Done()

	defer func() {
		doneCh <- struct{}{}
	}()

	{
		const gracePeriod = 10 * time.Second

		ctx, cancel := context.WithTimeout(context.Background(), gracePeriod)
		defer cancel()

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
