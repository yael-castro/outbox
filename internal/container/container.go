package container

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"sync"
)

type Container interface {
	Inject(context.Context, any) error
	Close(context.Context) error
}

type container struct {
	db  *sql.DB
	mux sync.Mutex
}

func (c *container) Inject(ctx context.Context, a any) error {
	switch a := a.(type) {
	case **sql.DB:
		return c.injectDB(ctx, a)
	}

	return fmt.Errorf("type \"%T\" is not supported", a)
}

func (c *container) Close(_ context.Context) error {
	return c.db.Close()
}

func (c *container) injectDB(ctx context.Context, db **sql.DB) (err error) {
	err = c.initDB(ctx)
	if err != nil {
		return
	}

	*db = c.db
	return err
}

func (c *container) initDB(ctx context.Context) (err error) {
	c.mux.Lock()
	defer c.mux.Unlock()

	if c.db != nil {
		return
	}

	const sqlDSN = "SQL_DSN"
	dsn := os.Getenv(sqlDSN)
	if len(dsn) == 0 {
		err = fmt.Errorf("missing '%s' environment variable", sqlDSN)
		return
	}

	var newDB *sql.DB

	const driverName = "postgres"
	newDB, err = sql.Open(driverName, dsn)
	if err != nil {
		return
	}

	err = newDB.PingContext(ctx)
	if err != nil {
		return
	}

	c.db = new(sql.DB)
	c.db = newDB
	return
}
