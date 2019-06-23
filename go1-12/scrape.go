package main

import (
	"log"

	gq "github.com/PuerkitoBio/goquery"
)

func main() {
	log.Print("start scrape")

	url := "https://godoc.org/github.com/PuerkitoBio/goquery"
	doc, err := gq.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	sel := doc.Find("h4#NewDocument")
	log.Print(sel.Text())

	log.Print("end scrape")
}
