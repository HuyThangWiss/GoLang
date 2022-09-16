package ControllersJwtB1

import (
	"b1/JWT/B1/ConnectJwtB1"
	"b1/JWT/B1/InforJwtB1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterUser(c *gin.Context) {
	var user InforJwtB1.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error ": err.Error()})
		c.Abort()
		return
	}
	if err := user.HashPassword(user.Userpassword); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error ": err.Error()})
		c.Abort()
		return
	}
	record := ConnectJwtB1.DB.Create(&user)

	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error ": record.Error.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{"User ": user.Username, "Password ": user.Userpassword, "PhoneNumber :": user.Phonenumber})
}
