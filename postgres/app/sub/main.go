package main

import (
	"context"
	"log"
	"os"

	// https://godoc.org/github.com/jackc/pgx
	"github.com/jackc/pgx/v4"
)

func main() {
	//DBDSN := "host=pgdb user=postgres port=5432"

	// https://github.com/jackc/pgx#example-usage
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_DSN"))
	if err != nil {
		panic(err)
	}

	var rslt int
	err = conn.QueryRow(context.Background(), "SELECT 1").Scan(&rslt)
	if err != nil {
		panic(err)
	}
	log.Println(rslt)

	listen := "LISTEN testpubsub"
	log.Println(listen + " start")

	// https://godoc.org/github.com/jackc/pgx#hdr-Listen_and_Notify
	_, err = conn.Exec(context.Background(), listen)
	if err != nil {
		panic(err)
	}

	log.Println("endless...")
	for {
		// err != nil の時に？ -> ではなかった
		notification, err := conn.WaitForNotification(context.Background())
		if err != nil {
			panic(err)
		}

		log.Println("catched notify!!" + " by " + os.Getenv("APP_NUMBER"))
		log.Printf("-notification-\n%+v\n", notification)
		// => &{PID:54 Channel:testpubsub Payload:nnnnnotify}

		log.Println("--Channel--")
		log.Println(notification.Channel)

		log.Println("--Payload--")
		log.Println(notification.Payload)
	}
}
