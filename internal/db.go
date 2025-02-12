package internal

import (
	"context"
	"os"

	_ "github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var urlExample1 string = os.Getenv("POSTGRES_URL")

var database *pgxpool.Pool

func InitDBPool() error {
	config, err := pgxpool.ParseConfig(urlExample1)
	if err != nil {
		// ...
	}
	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		// do something with every new connection
		return nil
	}
	database, err = pgxpool.NewWithConfig(context.Background(), config)
	return err
}

func DB() (*pgxpool.Pool, error) {
	err := database.Ping(context.Background())
	if err != nil {
		return nil, err
	} else {
		return database, nil
	}
}
