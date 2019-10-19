package main

import (
	"context"
	"log"
	"os"
	"time"

	// https://godoc.org/github.com/jackc/pgx
	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := connect(os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatalln(err)
	}

	// app/subがLISTENするのを待つため...
	time.Sleep(10 * time.Second)

	if err := notify(conn); err != nil {
		log.Fatalln(err)
	}
}

func connect(dsn string) (*pgx.Conn, error) {
	// https://github.com/jackc/pgx#example-usage
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	var rslt int
	err = conn.QueryRow(context.Background(), "SELECT 1").Scan(&rslt)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func notify(conn *pgx.Conn) error {
	notify := "NOTIFY testpubsub, 'nnnnnotify'"
	log.Println(notify)

	_, err := conn.Exec(context.Background(), notify)
	if err != nil {
		return err
	}

	return nil
}
