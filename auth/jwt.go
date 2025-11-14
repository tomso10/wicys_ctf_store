package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/config"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/data"
	"gitlab.ritsec.cloud/competitions/ists-2023/store/ent"
)

type Claims struct {
	Username string
	jwt.StandardClaims
}
type User struct {
	Username   string
	UserID     int
	Team       string
	Balance    int
	Expiration int
}

var (
	jwtKey []byte
)

var (
	ErrUnauthorized error = errors.New("unauthorized")
	ErrInvalidToken error = errors.New("invalid token")
)

func init() {
	jwtKey = []byte(config.JWTKey)
}

func GenerateJWT(username string) (string, int, error) {
	expiration := time.Now().Add(time.Duration(config.Timeout) * time.Minute)

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(jwtKey)

	return tokenStr, int(expiration.Unix()), err
}

func tokenParse(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return "", ErrInvalidToken
	}
	return jwtKey, nil
}

func ParseJWT(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, tokenParse)
	return token, claims, err
}

func Parse(ctx *gin.Context) (*ent.User, error) {
	tokenString, err := ctx.Cookie("auth")
	if err != nil {
		return nil, ErrUnauthorized
	}

	_, claims, err := ParseJWT(tokenString)
	if err != nil {
		return nil, err
	}

	user, err := data.User.Get(claims.Username)
	if err != nil {
		return nil, err
	}

	return user, err
}
