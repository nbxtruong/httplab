package main

import (
	"net/http"

	"github.com/graphql-go/handler"
)

func init()  {
	importJSONFromFile("./mock/beast.json", &BeastList)
}

func main() {
	h := handler.New(&handler.Config{
		Schema:   &BeastSchema,
		Pretty:   true,
		GraphiQL: false,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}
