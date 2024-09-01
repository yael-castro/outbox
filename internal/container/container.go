package container

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func inject(ctx context.Context, a any) error {
	switch a := a.(type) {
	case *sql.DB:
		return injectSQLDB(ctx, a)
	}

	return fmt.Errorf("type \"%T\" is not supported", a)
}

func injectSQLDB(ctx context.Context, db *sql.DB) (err error) {
	newDB, err := SingleDB(ctx)
	if err != nil {
		return
	}

	*db = *newDB
	return
}
