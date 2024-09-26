package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

const connString = "postgresql://localhost:5432/kemmer"

func main() {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var greeting string
	err = conn.QueryRow(context.Background(), "SELECT 'hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "query row failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
