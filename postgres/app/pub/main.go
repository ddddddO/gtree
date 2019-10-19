package main

import (
	"context"
	"fmt"

	// https://godoc.org/github.com/jackc/pgx
	"github.com/jackc/pgx/v4"
)

func main() {
	fmt.Println("vim-go")

	//DBDSN := "host=172.18.0.2 user=postgres port=5432"
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

	fmt.Println(rslt)
}
