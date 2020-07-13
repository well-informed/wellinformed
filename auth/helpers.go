package auth

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/well-informed/wellinformed/graph/model"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword - Hashes Passwords
func HashPassword(password string) (hashPassword string, err error) {
	bytePassword := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// u.Password = string(passwordHash)

	return string(passwordHash), nil
}

// GenToken - Generates JWT Tokens
func GenToken(userID int64, expiredAt time.Time, env string) (*model.AuthToken, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expiredAt.Unix(),
		Id:        strconv.FormatInt(userID, 10),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "wellinformed",
	})

	accessToken, err := token.SignedString([]byte(os.Getenv(env)))
	if err != nil {
		return nil, err
	}

	return &model.AuthToken{
		AccessToken: accessToken,
		ExpiredAt:   expiredAt,
	}, nil
}

func GenAccessToken(userID int64) (*model.AuthToken, error) {
	return GenToken(userID, time.Now().Add(time.Minute*30), "JWT_ACCESS_SECRET") // 30 minutes
}

func GenRefreshToken(userdID int64) (*model.AuthToken, error) {
	return GenToken(userdID, time.Now().Add(time.Hour*24*7), "JWT_REFRESH_SECRET") // a week
}

// ComparePassword - Compares Passwords
func ComparePassword(password string, existingPassword string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(existingPassword)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
