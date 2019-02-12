package main

/*
	参考
	https://qiita.com/katsumic/items/a7984afca2d4522f60ac

	TODO:
	・既存のシートを編集する処理追加
	・スプレッドシート編集関数細分化(editSpreadSheet)
*/

import (
	"io/ioutil"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

func main() {
	// コンテキストを生成
	ctx := context.Background() // NOTE:こいつの用途

	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %+v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(
		b,
		"https://www.googleapis.com/auth/spreadsheets",
		"https://www.googleapis.com/auth/drive.file",
	)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %+v", err)
	}

	client := getClient(config, ctx)

	service, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %+v", err)
	}

	createResponse, err := genSpreadSheet(service, ctx)
	if err != nil {
		log.Fatal(err)
	}

	// NOTE:
	m := map[string][]interface{}{}
	m["code"] = []interface{}{"a", "b", "c", "d", "e"}

	if err := editSpreadSheet(service, createResponse, m); err != nil {
		log.Fatal(err)
	}

}
