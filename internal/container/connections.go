package container

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"sync"
)

var (
	onceSQL  sync.Once
	singleDB *sql.DB
)

func SingleDB(ctx context.Context) (*sql.DB, error) {
	var err error

	onceSQL.Do(func() {
		const sqlDSN = "SQL_DSN"

		dsn := os.Getenv(sqlDSN)
		if len(dsn) == 0 {
			err = fmt.Errorf("missing '%s' environment variable", sqlDSN)
			return
		}

		const driverName = "postgres"
		var newDB *sql.DB

		newDB, err = sql.Open(driverName, dsn)
		if err != nil {
			return
		}

		err = newDB.PingContext(ctx)
		if err != nil {
			return
		}

		singleDB = new(sql.DB)
		*singleDB = *newDB // TODO: suppress warning
	})

	return singleDB, err
}
