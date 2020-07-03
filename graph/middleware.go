package graph

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/pkg/errors"
	"github.com/well-informed/wellinformed"
	"github.com/well-informed/wellinformed/graph/model"
)

const CurrentUserKey = "currentUser"

func AuthMiddleware(db wellinformed.Persistor) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := parseToken(r)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)

			if !ok || !token.Valid {
				next.ServeHTTP(w, r)
				return
			}

			user, err := db.GetUserById(claims["jti"].(string))
			// fmt.Println(user)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), CurrentUserKey, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

var authHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixFromToken,
}

func stripBearerPrefixFromToken(token string) (string, error) {
	bearer := "BEARER"

	if len(token) > len(bearer) && strings.ToUpper(token[0:len(bearer)]) == bearer {
		return token[len(bearer)+1:], nil
	}

	return token, nil
}

var authExtractor = &request.MultiExtractor{
	authHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

func parseToken(r *http.Request) (*jwt.Token, error) {
	jwtToken, err := request.ParseFromRequest(r, authExtractor, func(token *jwt.Token) (interface{}, error) {
		t := []byte(os.Getenv("JWT_SECRET"))
		return t, nil
	})

	return jwtToken, errors.Wrap(err, "parseToken error: ")
}

func GetCurrentUserFromCTX(ctx context.Context) (*model.User, error) {
	errNoUserInContext := errors.New("no user in context")

	if ctx.Value(CurrentUserKey) == nil {
		return nil, errNoUserInContext
	}

	fmt.Println(ctx.Value(CurrentUserKey))
	user, ok := ctx.Value(CurrentUserKey).(model.User)
	if !ok {
		return nil, errNoUserInContext
	}

	return &user, nil
}
