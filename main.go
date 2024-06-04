package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	generated "github.com/yofriadi/backend-vanilla/generated/graphql"
	"github.com/yofriadi/backend-vanilla/generated/merchandise/design/v1/designv1connect"
	"github.com/yofriadi/backend-vanilla/graphql/resolver"
	designv1 "github.com/yofriadi/backend-vanilla/rpc/merchandise/design/v1"
)

func main() {
	mux := http.NewServeMux()

	graphqlServer := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}),
	)
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", graphqlServer)

	designServerV1 := &designv1.DesignServer{}
	path, handler := designv1connect.NewDesignServiceHandler(designServerV1)
	mux.Handle(path, handler)

	log.Fatal(http.ListenAndServe(":8080", h2c.NewHandler(mux, &http2.Server{})))
}
