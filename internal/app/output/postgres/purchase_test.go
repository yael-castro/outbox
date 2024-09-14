//go:build tests && http

package postgres_test

import (
	"context"
	"database/sql"
	"github.com/yael-castro/outbox/internal/app/business"
	"github.com/yael-castro/outbox/internal/app/output/postgres"
	"github.com/yael-castro/outbox/internal/container"
	"log"
	"os"
	"strconv"
	"testing"
)

func TestPurchaseSaver_SavePurchase(t *testing.T) {
	cases := [...]struct {
		ctx      context.Context
		purchase *business.Purchase
	}{
		{
			ctx: context.Background(),
			purchase: &business.Purchase{
				OrderID: 1_000,
			},
		},
	}

	ctx := context.Background()

	// Establishing connection with DB
	var db *sql.DB

	c := container.New()

	err := c.Inject(ctx, &db)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		_ = c.Close(context.Background())
	})

	topic := os.Getenv("KAFKA_SERVERS")
	if len(topic) == 0 {
		t.Fatal("Missing environment variable!")
	}

	errLogger := log.New(os.Stderr, "[ERROR] ", log.LstdFlags)
	saver := postgres.NewPurchaseSaver(postgres.PurchaseSaverConfig{
		DB:     db,
		Topic:  topic,
		Logger: errLogger,
	})

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ctx := c.ctx

			err := saver.SavePurchase(ctx, c.purchase)
			if err != nil {
				t.Fatal(err)
			}

			t.Cleanup(func() {
				// TODO: Delete created records
			})

			t.Logf("Purchase: %+v", c.purchase)
		})
	}
}
