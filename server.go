package main

import (
	"github.com/Raven57/belajargraphql/graph/postgre"
	"github.com/go-pg/pg/v10"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Raven57/belajargraphql/graph"
	"github.com/Raven57/belajargraphql/graph/generated"
)

const defaultPort = "8888"

func main() {

	pgDB :=postgre.New(&pg.Options{
		Addr:	":5432",
		User: "postgres",
		Password: "xperianeo5%",
		Database: "test",
	})

	pgDB.AddQueryHook(postgre.DBLogger{})

	defer pgDB.Close()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		UsersRepo: postgre.UsersRepo{DB: pgDB},
		PremiumtypesRepo: postgre.PremiumtypesRepo{DB: pgDB},
		CommentsRepo :postgre.CommentsRepo{DB: pgDB},
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
