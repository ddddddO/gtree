package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/pkg/errors"
)

func main() {
	// neo4j
	dsn := "neo4j://localhost:7687"
	driver, err := neo4j.NewDriver(dsn, neo4j.BasicAuth("username", "password", ""))
	if err != nil {
		log.Fatal(err)
	}
	defer driver.Close()

	// graphql
	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// TODO: ここにneo4jへアクセスする処理をいくのがよいのかどうか
				if err := matchItem(driver); err != nil {
					return "", err
				}
				return "world", nil
			},
		},
		"morning": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	log.Println("start")

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	mux := http.NewServeMux()
	mux.Handle("/graphql", h)
	mux.HandleFunc("/health", health(driver))
	mux.HandleFunc("/debugNeo4j", debug(driver))
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// graceful shutdown...
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	<-sig

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

}

func health(driver neo4j.Driver) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if err := matchItem(driver); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf(`{"msg":"%v"}`, err)))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"msg":"health ok"}`))
	}
}

func matchItem(driver neo4j.Driver) error {
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	getFn := func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run("MATCH (i:Item) RETURN i LIMIT 1", nil)
		if err != nil {
			return nil, err
		}

		if !result.Next() {
			return nil, errors.New("no item")
		}
		if err := result.Err(); err != nil {
			return nil, err
		}
		return nil, nil
	}

	_, err := session.ReadTransaction(getFn)
	if err != nil {
		return err
	}

	return nil
}

func debug(driver neo4j.Driver) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		item, err := insertItem(driver)
		if err != nil {
			panic(err)
		}
		log.Printf("%+v\n", item)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("%+v", item)))
	}
}

func insertItem(driver neo4j.Driver) (*Item, error) {
	// Sessions are short-lived, cheap to create and NOT thread safe. Typically create one or more sessions
	// per request in your web application. Make sure to call Close on the session when done.
	// For multi-database support, set sessionConfig.DatabaseName to requested database
	// Session config will default to write mode, if only reads are to be used configure session for
	// read mode.
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, err := session.WriteTransaction(createItemFn)
	if err != nil {
		return nil, err
	}
	return result.(*Item), nil
}

func createItemFn(tx neo4j.Transaction) (interface{}, error) {
	records, err := tx.Run("CREATE (n:Item { id: $id, name: $name }) RETURN n.id, n.name", map[string]interface{}{
		"id":   1,
		"name": "Item 1",
	})
	// In face of driver native errors, make sure to return them directly.
	// Depending on the error, the driver may try to execute the function again.
	if err != nil {
		return nil, err
	}
	record, err := records.Single()
	if err != nil {
		return nil, err
	}
	// You can also retrieve values by name, with e.g. `id, found := record.Get("n.id")`
	return &Item{
		Id:   record.Values[0].(int64),
		Name: record.Values[1].(string),
	}, nil
}

type Item struct {
	Id   int64
	Name string
}
