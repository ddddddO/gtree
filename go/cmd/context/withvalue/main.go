package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "ID", "1234")
	ctx = context.WithValue(ctx, "Name", "XXXX")

	log.Println("ID:", ctx.Value("ID"))
	log.Println("Name:", ctx.Value("Name"))
}