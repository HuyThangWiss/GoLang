package TokenJwtB3

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type JWTService interface {
	GenerateToken(Masv string, Passwordsv string) string
	ValidateToken(tokenString string) (*jwt.Token, error)
}

// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
	Masv       string `json:"masv" binding:"required"`
	Passwordsv string `json:"passwordsv" binding:"required"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
}

func NewJWTService() *jwtService {
	return &jwtService{
		secretKey: string("mabuu"),
	}
}

func (jwtSrv *jwtService) GenerateToken(Masv string, Passwordsv string) string {

	// Set custom and standard claims
	claims := &jwtCustomClaims{
		Masv,
		Passwordsv,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),

			IssuedAt: time.Now().Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token using the secret signing key
	t, err := token.SignedString([]byte(jwtSrv.secretKey))
	if err != nil {
		//	panic(err)
		log.Println(err)
	}
	return t
}

func (jwtSrv *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret signing key
		return []byte(jwtSrv.secretKey), nil
	})
}
