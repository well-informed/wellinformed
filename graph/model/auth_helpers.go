package model

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword - Hashes Passwords
func (u *User) HashPassword(password string) error {
	bytePassword := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(passwordHash)

	return nil
}

// GenToken - Generates JWT Tokens
func (u *User) GenToken(expiredAt time.Time, env string) (*AuthToken, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expiredAt.Unix(),
		Id:        strconv.FormatInt(u.ID, 10),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "wellinformed",
	})

	accessToken, err := token.SignedString([]byte(os.Getenv(env)))
	if err != nil {
		return nil, err
	}

	return &AuthToken{
		AccessToken: accessToken,
		ExpiredAt:   expiredAt,
	}, nil
}

func (u *User) GenAccessToken() (*AuthToken, error) {
	return u.GenToken(time.Now().Add(time.Minute*30), "JWT_ACCESS_SECRET") // 30 minutes
}

func (u *User) GenRefreshToken() (*AuthToken, error) {
	return u.GenToken(time.Now().Add(time.Hour*24*7), "JWT_REFRESH_SECRET") // a week
}

// ComparePassword - Compares Passwords
func (u *User) ComparePassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
