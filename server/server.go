package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed/database"
	"github.com/well-informed/wellinformed/graph"
	"github.com/well-informed/wellinformed/graph/generated"
	"github.com/well-informed/wellinformed/rss"
)

const defaultPort = "8080"

func main() {
	log.SetLevel(log.DebugLevel)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	resolver := &graph.Resolver{
		DB:  database.NewDB(),
		RSS: rss.NewRSS(),
	}

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	// router.Use(middleware.RequestID)
	// router.Use(middleware.Logger)
	router.Use(graph.AuthMiddleware(resolver.DB))

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
