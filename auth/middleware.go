package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
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
			token, err := parseToken(r)
			// _, cookieErr := r.Cookie("jid")

			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)

			if !ok || !token.Valid {
				next.ServeHTTP(w, r)
				return
			}

			id, err := strconv.ParseInt(claims["jti"].(string), 10, 64)
			if err != nil {
				log.Error("could not parse claim into int. err: ", err)
				next.ServeHTTP(w, r)
				return
			}

			user, err := db.GetUserByID(id)
			if err != nil || user == nil {
				next.ServeHTTP(w, r)
				return
			}

			// set the current user in context
			ctx := context.WithValue(r.Context(), CurrentUserKey, user)
			log.Trace("setting user context value to: ", user)
			refreshToken, err := GenRefreshToken(user.ID)
			if err != nil {
				log.Errorf("Error refreshing user token: %s\n", err.Error())
				next.ServeHTTP(w, r)
				return
			}
			http.SetCookie(w, &http.Cookie{
				Name:     "jid",
				Value:    refreshToken.AccessToken,
				HttpOnly: true,
			})
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
		t := []byte(os.Getenv("JWT_ACCESS_SECRET"))
		return t, nil
	})
	return jwtToken, errors.Wrap(err, "parseToken error: ")
}

func GetCurrentUserFromCTX(ctx context.Context) (*model.User, error) {
	errNoUserInContext := errors.New("no user in context")

	if ctx.Value(CurrentUserKey) == nil {
		log.Error("current user key is empty.", errNoUserInContext)
		return nil, errNoUserInContext
	}

	fmt.Println(ctx.Value(CurrentUserKey))
	user, ok := ctx.Value(CurrentUserKey).(*model.User)
	if !ok || user == nil {
		log.Error("could not parse current user object.", errNoUserInContext)
		return nil, errNoUserInContext
	}
	log.Debug("Got user from context: ", user)
	return user, nil
}

func RefreshToken(db wellinformed.Persistor) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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

		id, err := strconv.ParseInt(claims["jti"].(string), 10, 64)
		if err != nil {
			log.Error("could not parse claim into int. err: ", err)
			w.Write(errJsonRes)
			return
		}
		user, err := db.GetUserByID(id)

		log.Printf("user: %v", user)

		if err != nil || user == nil {
			log.Printf("err getting user from db: %v", err)
			w.Write(errJsonRes)
			return
		}

		refreshToken, err := GenRefreshToken(user.ID)
		accessToken, err := GenAccessToken(user.ID)

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
