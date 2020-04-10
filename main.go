package main

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"log"
	"net/http"
)

// TODO: Schema
// TODO: Model
type query struct{}

// TODO: Resolver
func main() {
	schema := graphql.MustParseSchema(Schema, &query{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	// TODO: init model
	// TODO: graphiql
	activateGraphiql()

	// Run
	log.Println("Server ready at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
