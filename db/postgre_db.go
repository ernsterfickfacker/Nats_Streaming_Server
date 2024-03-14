package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Conn *pgxpool.Pool

func DBConnection() {
	var err error
	connStr := "postgres://postgres:admin@localhost:5432/postgres?sslmode=disable"
	Conn, err = pgxpool.New(context.Background(),
		connStr)
	if err != nil {
		fmt.Println("Error connecting to the database: %v\n", err)
		panic("Error connecting to the database")
		os.Exit(1)
	}
}
