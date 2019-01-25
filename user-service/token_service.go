package main

import (
	"time"

	pb "github.com/shooshpanov/microservices-project/user-service/proto/user"
	"github.com/dgrijalva/jwt-go"

)

var (
	// salt key
	key := []byte("mySurepSecretKey")
)

type CustomClaims struct {
	User *pb.User
	jwt.StandartClaims
}

type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type TokenService struct {
	repo Repository
}

func (srv *TokenService) Decode(tokenString string) (*CustomClaims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
}

func (srv *TokenService) Encode(user *pb.User) (string, error) {

	expireToken := time.Now().Add(time.Hour * 72).Unix()

	// Create the Claims
	claims := CustomClaims{
		user,
		jwt.StandartClaims{
			ExpireAt: expireToken,
			Issuer: "go.micro.srv.user",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	return token.SignedString(key)
}