package TokenJwtB3

//goland:noinspection ALL
import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	Codeid string `json:"codeid"`
	Pass   string `json:"pass"`
	jwt.StandardClaims
}

func GenerateJWTauth(Codeid string, Pass string) (tokenString string, err error) {

	expiraion := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Codeid: Codeid,
		Pass:   Pass,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiraion.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SigningString()
	return
}
func ValidateTokenauth(signedToken string) (err error) {
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
		err = errors.New("Couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("Token expired")
		return
	}
	return
}
