package MiddwareJwtB3

import (
	"b1/JWT/B3/TokenJwtB3"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("User ")
		if tokenString == "" {
			c.JSON(401, gin.H{"erroe ": "request does not contain an access token"})
			c.Abort()
			return
		}
		err := TokenJwtB3.ValidateTokenauth(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error ": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
