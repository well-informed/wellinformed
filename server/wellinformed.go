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
	"github.com/well-informed/wellinformed/graph"
	"github.com/well-informed/wellinformed/graph/generated"
	"github.com/well-informed/wellinformed/rss"
	"github.com/well-informed/wellinformed/subscriber"
	"github.com/well-informed/wellinformed/user"
	"github.com/well-informed/wellinformed/userFeed"
)

const defaultPort = "8080"

func main() {

	conf := wellinformed.GetConfig()

	log.SetLevel(conf.LogLevel)

	db := database.NewDB(conf)
	rss := rss.NewRSS()
	sub, err := subscriber.NewSubscriber(rss, db)
	if err != nil {
		log.Fatal("couldn't initialize new subscriber properly")
	}
	feedService := userFeed.NewFeedService(db)

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
		AllowedOrigins:   []string{"http://alpha.edyn.me", "https://alpha.edyn.me", "http://api.edyn.me/", "https://api.edyn.me/"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
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

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", conf.ServerPort)
	log.Fatal(http.ListenAndServe(":"+conf.ServerPort, router))
}
