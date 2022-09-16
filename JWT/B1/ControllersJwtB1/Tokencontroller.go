package ControllersJwtB1

import "C"
import (
	"b1/JWT/B1/ConnectJwtB1"
	"b1/JWT/B1/InforJwtB1"
	"b1/JWT/B1/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type TokenRequestJwtb1 struct {
	Username     string `json:"user_name"`
	Userpassword string `json:"user_password"`
}

func GenerateToken(c *gin.Context) {
	var request TokenRequestJwtb1
	var user InforJwtB1.Users

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		c.Abort()
		return
	}
	record := ConnectJwtB1.DB.Where("Username =  ?", request.Username).First(&user)

	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error ": record.Error.Error()})
		c.Abort()
		return
	}
	credentialError := user.CheckPassword(request.Userpassword)
	if credentialError != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Errorn :": "Invalid credentials"})
		c.Abort()
		return
	}
	tokenString, err := auth.GenerateJWTB1(user.Username, user.Phonenumber)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func SelectJwtB1(c *gin.Context) {
	ConnectJwtB1.ConnectJwtB1()
	var arr []InforJwtB1.Users
	ConnectJwtB1.DB.Find(&arr)
	c.JSON(http.StatusOK, gin.H{"Data ": arr})
}

func EnCode(password string) (string, error) {

	var passwordBytes = []byte(password)

	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	return string(hashedPasswordBytes), err
}

func InsertJwtB1(c *gin.Context) {
	ConnectJwtB1.ConnectJwtB1()
	var Input InforJwtB1.Users
	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error ": err.Error()})
		return
	}
	hash, _ := EnCode(Input.Userpassword)
	human := InforJwtB1.Users{
		Macc:         Input.Macc,
		Username:     Input.Username,
		Userpassword: hash,
		Address:      Input.Address,
		Phonenumber:  Input.Phonenumber,
	}
	ConnectJwtB1.DB.Create(&human)
	c.JSON(http.StatusOK, gin.H{"Data : ": human})
}
