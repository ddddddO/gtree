package main

import (
	"log"

	"github.com/ddddddO/work/go-gui/v2/db"
	"github.com/ddddddO/work/go-gui/v2/ui"
)

func main() {
	sqlite, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlite.CloseSQLite()

	log.Println("start")

	ui.Run(sqlite)
}
