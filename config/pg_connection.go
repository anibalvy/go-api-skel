package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)


func PGconn() (*pgx.Conn, error) {
    db_url := Conf["db_url"].(string)
    conn, err := pgx.Connect(context.Background(), db_url)
    if err != nil {
        fmt.Printf("PGconn - Unable to connect to DB - err: %v", err)
    }

    return conn, err
}
