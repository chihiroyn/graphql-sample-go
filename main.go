package main

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"log"
	"net/http"
)

// Schema
const Schema = `
type Vegetable {
    name: String!
    price: Int!
    image: String
}
type Query {
    vegetable(name: String!): Vegetable
}
schema {
    query: Query
}
`

type Vegetable struct {
	name  string
	price int
	image *string
}

var vegetables map[string]Vegetable

// Utils
func strPtr(str string) *string {
	return &str
}


// TODO: Model
type query struct{}

func main() {
	schema := graphql.MustParseSchema(Schema, &query{})
	http.Handle("/query", &relay.Handler{Schema: schema})

	// init model
	vegetables = map[string]Vegetable{
		"tomato": Vegetable{name: "Tomato", price: 100, image: strPtr("https://picsum.photos/id/152/100/100")},
		"potato": Vegetable{name: "Potato", price: 50, image: strPtr("https://picsum.photos/id/159/100/100")},
		"corn":   Vegetable{name: "Corn", price: 200},
	}

	// TODO: graphiql
	activateGraphiql()

	// Run
	log.Println("Server ready at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
