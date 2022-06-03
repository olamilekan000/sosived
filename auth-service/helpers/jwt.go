package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/olamilekan000/sosivio-app/auth-service/repository"
)

type JWTStruct struct{}

func (j *JWTStruct) VerifyJWTToken(jwtToken string) (*repository.AuthResponse, error) {
	parsedToken, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if err, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Printf(err.Name)
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	if parsedToken.Valid {
		claims, _ := parsedToken.Claims.(jwt.MapClaims)
		iss := claims["iss"]
		return &repository.AuthResponse{User: iss}, nil
	}
	return nil, err
}

func (j *JWTStruct) CreateToken(username string) (string, error) {
	var err error

	atClaims := jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		Issuer:    username,
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}
	return token, nil
}
