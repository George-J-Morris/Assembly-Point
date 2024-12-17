package internal

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type suAccount struct {
	usernamename string
	password     string
}

var urlExample1 string = ""

var database *pgxpool.Pool

func GetSuperUserByName(username string, connPool *pgxpool.Pool) suAccount {

	var resp suAccount

	query := "SELECT * FROM users WHERE username = $1"

	err := connPool.Ping(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	row := connPool.QueryRow(context.Background(), query, username)

	err = row.Scan(&resp.usernamename, &resp.password)

	if err != nil {
		fmt.Println(err)
	}

	return resp
}

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
