package handler

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type graphqlHandler struct {
	driver neo4j.Driver
}

func NewGraphqlHandler(driver neo4j.Driver) *graphqlHandler {
	return &graphqlHandler{
		driver: driver,
	}
}

func (gh *graphqlHandler) Handler() (*handler.Handler, error) {
	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// TODO: ここにneo4jへアクセスする処理をいくのがよいのかどうか
				if err := matchItem(gh.driver); err != nil {
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
		return nil, err
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	return h, nil
}
