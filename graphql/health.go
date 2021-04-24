package main

import (
	"fmt"
	"net/http"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/pkg/errors"
)

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
