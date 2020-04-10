package main

import (
	"github.com/friendsofgo/graphiql"
	"net/http"
)

func activateGraphiql() {
	// First argument must be same as graphql handler path
	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/query")
	if err != nil {
		panic(err)
	}
	http.Handle("/", graphiqlHandler)
}
