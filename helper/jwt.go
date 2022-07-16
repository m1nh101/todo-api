package helper

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/m1ngi/todo-api/model"
)

var (
	JWT_PUBLIC_KEY = os.Getenv("JWT_PUBLIC_KEY")
)

func GenerateJwtToken(user *model.User) string {
	expiredTokenTime := time.Now().UTC().Add(time.Hour * 24).Unix()

	tokenDescription := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf":            time.Now().UTC().Unix(),
		"iat":            time.Now().UTC().Unix(),
		"exp":            expiredTokenTime,
		"id":             user.Username,
		"nameIdentifier": user.Id.String(),
	})

	token, err := tokenDescription.SignedString([]byte(JWT_PUBLIC_KEY))

	if err != nil {
		log.Fatal(err)
	}

	return token
}

func ValidateJwtToken(token string) (interface{}, error) {
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return []byte(JWT_PUBLIC_KEY), nil
	})
	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, fmt.Errorf("validate: invalid")
	}

	return claims["nameIdentifier"], nil
}
