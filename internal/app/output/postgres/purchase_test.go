//go:build tests

package postgres

import (
	"context"
	"database/sql"
	"github.com/yael-castro/outbox/internal/app/business"
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
	var db sql.DB

	err := container.Inject(ctx, &db)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		_ = db.Close()
	})

	errLogger := log.New(os.Stderr, "[ERROR] ", log.LstdFlags)
	saver := NewPurchaseSaver(&db, errLogger)

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := saver.SavePurchase(c.ctx, c.purchase)
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
