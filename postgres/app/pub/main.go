package main

import (
	"context"
	"log"
	"time"

	// https://godoc.org/github.com/jackc/pgx
	"github.com/jackc/pgx/v4"
)

func main() {
	DBDSN := "host=pgdb user=postgres port=5432"

	// https://github.com/jackc/pgx#example-usage
	conn, err := pgx.Connect(context.Background(), DBDSN)
	if err != nil {
		panic(err)
	}

	var rslt int
	err = conn.QueryRow(context.Background(), "SELECT 1").Scan(&rslt)
	if err != nil {
		panic(err)
	}
	log.Println(rslt)

	// app/subがLISTENするのを待つため...
	time.Sleep(10 * time.Second)

	notify := "NOTIFY testpubsub, 'nnnnnotify'"
	log.Println(notify)

	_, err = conn.Exec(context.Background(), notify)
	if err != nil {
		panic(err)
	}
}
