package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	Username    string `json:"user_name"`
	Phonenumber string `json:"phonenumber"`
	jwt.StandardClaims
}

func GenerateJWTB1(PhoneNumber string, UserName string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Username:    UserName,
		Phonenumber: PhoneNumber,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateTokenJWTB1(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
