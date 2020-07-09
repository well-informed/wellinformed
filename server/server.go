package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed/auth"
	"github.com/well-informed/wellinformed/database"
	"github.com/well-informed/wellinformed/graph"
	"github.com/well-informed/wellinformed/graph/generated"
	"github.com/well-informed/wellinformed/graph/model"
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
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	// router.Use(middleware.RequestID)
	// router.Use(middleware.Logger)
	router.Use(auth.AuthMiddleware(resolver.DB))
	router.Use(render.SetContentType(render.ContentTypeJSON))

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	router.Post("/refresh_token", func(w http.ResponseWriter, r *http.Request) {

		type JSONResponse struct {
			Ok          bool             `json:"ok"`
			AccessToken *model.AuthToken `json:"accessToken"`
		}

		w.Header().Set("Content-Type", "application/json")

		token, err := r.Cookie("jid")

		errJsonRes, _ := json.Marshal(JSONResponse{
			Ok:          false,
			AccessToken: nil,
		})

		if err != nil {
			log.Printf("err getting token from cookie: %v", err)
			w.Write(errJsonRes)
			return
		}

		claims, err := ValidateRefreshToken(token.Value)

		if err != nil {
			log.Printf("err ValidateRefreshToken: %v", err)
			w.Write(errJsonRes)
			return
		}

		user, err := resolver.DB.GetUserById(claims["jti"].(string))

		log.Printf("user: %v", user)

		if err != nil {
			log.Printf("err getting user from db: %v", err)
			w.Write(errJsonRes)
			return
		}

		refreshToken, err := user.GenRefreshToken()
		accessToken, err := user.GenAccessToken()

		// check token version maybe here?

		http.SetCookie(w, &http.Cookie{
			Name:     "jid",
			Value:    refreshToken.AccessToken,
			HttpOnly: true,
		})

		res := JSONResponse{
			Ok:          true,
			AccessToken: accessToken,
		}

		jsonres, err := json.Marshal(res)
		w.Write(jsonres)

	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// ValidateToken ValidateToken
func ValidateToken(value string, env string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(value, func(token *jwt.Token) (interface{}, error) {
		// Dont forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header)
		}

		return []byte(os.Getenv(env)), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
		return claims, nil
	} else {
		fmt.Println(err)
		return nil, err
	}
}

// ValidateRefreshToken ValidateRefreshToken
func ValidateRefreshToken(value string) (jwt.MapClaims, error) {
	return ValidateToken(value, "JWT_REFRESH_SECRET")
}
