package auth

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed"
	"github.com/well-informed/wellinformed/graph/model"
)

const CurrentUserKey = "currentUser"

var (
	ErrBadCredentials  = errors.New("email/password combination don't work")
	ErrUnauthenticated = errors.New("unauthenticated")
	ErrForbidden       = errors.New("unauthorized")
)

func AuthMiddleware(db wellinformed.Persistor) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Info("in the auth middleware")
			token, err := parseToken(r)
			_, cookieErr := r.Cookie("jid")

			if err != nil {
				log.Error("error with parseToken")
				next.ServeHTTP(w, r)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)

			if !ok || !token.Valid {
				log.Errorf("token not valid")
				next.ServeHTTP(w, r)
				return
			}

			user, err := db.GetUserById(claims["jti"].(string))
			// fmt.Println(user)
			if err != nil {
				log.Errorf("err getting user from token")
				next.ServeHTTP(w, r)
				return
			}

			// set the current user in context
			ctx := context.WithValue(r.Context(), CurrentUserKey, user)
			log.Infof("set ctx")
			if cookieErr != nil {
				refreshToken, err := user.GenRefreshToken()
				if err != nil {
					next.ServeHTTP(w, r)
					return
				}
				log.Errorf("refreshToken: %v", refreshToken.AccessToken)
				http.SetCookie(w, &http.Cookie{
					Name:     "jid",
					Value:    refreshToken.AccessToken,
					HttpOnly: true,
				})
			}
			log.Info("end of auth. ctx: ", ctx)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

var authHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixFromToken,
}

func stripPrefixFromValue(prefix string, value string) (string, error) {

	if len(value) > len(prefix) && strings.ToUpper(value[0:len(prefix)]) == prefix {
		return value[len(prefix)+1:], nil
	}

	return value, nil
}

func stripBearerPrefixFromToken(token string) (string, error) {
	return stripPrefixFromValue("BEARER", token)
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
		log.Error(errNoUserInContext)
		return nil, errNoUserInContext
	}

	fmt.Println(ctx.Value(CurrentUserKey))
	user, ok := ctx.Value(CurrentUserKey).(model.User)
	if !ok {
		log.Error(errNoUserInContext)
		return nil, errNoUserInContext
	}

	return &user, nil
}
