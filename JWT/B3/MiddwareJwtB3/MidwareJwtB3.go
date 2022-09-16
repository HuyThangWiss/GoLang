package MiddwareJwtB3

import (
	"b1/JWT/B3/TokenJwtB3"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		//const BEARER_SCHEMA = "Bearer "
		//authHeader := c.GetHeader("Authorization")
		//tokenString := authHeader[len(BEARER_SCHEMA):]
		tokenString := c.GetHeader("User")
		//		token, err := NewJWTService().ValidateToken(tokenString)
		token, err := TokenJwtB3.NewJWTService().ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Name]: ", claims["name"])
			log.Println("Claims[Admin]: ", claims["admin"])
			log.Println("Claims[ExpiresAt]: ", claims["exp"])

		} else {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
