package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed"
	"github.com/well-informed/wellinformed/auth"
	"github.com/well-informed/wellinformed/database"
	"github.com/well-informed/wellinformed/feed"
	"github.com/well-informed/wellinformed/graph"
	"github.com/well-informed/wellinformed/graph/generated"
	"github.com/well-informed/wellinformed/rss"
	"github.com/well-informed/wellinformed/subscriber"
	"github.com/well-informed/wellinformed/user"
)

const defaultPort = "8080"

func main() {
	conf := wellinformed.GetConfig()

	//Injecting database dependency so test version can be substituted as needed
	router, _ := initWellinformedApp(conf, database.NewDB(conf))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", conf.ServerPort)
	log.Fatal(http.ListenAndServe(":"+conf.ServerPort, router))
}

func initWellinformedApp(conf wellinformed.Config, db database.DB) (*chi.Mux, *graph.Resolver) {
	log.SetLevel(conf.LogLevel)

	rss := rss.NewRSS()
	sub, err := subscriber.NewSubscriber(rss, db)
	if err != nil {
		log.Fatal("couldn't initialize new subscriber properly")
	}
	feedService := feed.NewFeedService(db)

	resolver := &graph.Resolver{
		DB:          db,
		RSS:         rss,
		Sub:         sub,
		FeedService: feedService,
		UserService: user.NewUserService(db),
	}

	router := chi.NewRouter()

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   conf.CORSOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "FETCH"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
	}))

	// router.Use(middleware.RequestID)
	// router.Use(middleware.Logger)
	router.Use(auth.AuthMiddleware(resolver.DB))
	router.Use(render.SetContentType(render.ContentTypeJSON))

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	router.Post("/refresh_token", auth.RefreshToken(resolver.DB))
	return router, resolver
}
