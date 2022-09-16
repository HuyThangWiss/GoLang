package ControllerJwtB3

import (
	"b1/JWT/B2/JwtToken"
	"b1/JWT/B3/ConnectJwtB3"
	"b1/JWT/B3/InforJwtB3"
	"b1/JWT/B3/TokenJwtB3"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Encode(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func Select(c *gin.Context) {
	ConnectJwtB3.ConnectJwtB3()

	var arr []InforJwtB3.Staffs
	ConnectJwtB3.DB.Find(&arr)

	c.JSON(http.StatusOK, gin.H{"Data ": arr})
}

func Insert(c *gin.Context) {
	ConnectJwtB3.ConnectJwtB3()
	var Input InforJwtB3.Staffs
	//InforJwtB2
	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}

	hash, _ := Encode(Input.Pass)

	Admin := InforJwtB3.Staffs{
		Codeid: Input.Codeid,
		Fname:  Input.Fname,
		Pass:   hash,
		Adress: Input.Adress,
	}
	ConnectJwtB3.DB.Create(Admin)

	c.JSON(http.StatusOK, gin.H{"Data ": Admin})

}

func Logint(c *gin.Context) {
	ConnectJwtB3.ConnectJwtB3()
	var Input InforJwtB3.Staffs

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error : ": err.Error()})
		return
	}
	var admin InforJwtB3.Staffs
	result := ConnectJwtB3.DB.Where(&InforJwtB3.Staffs{
		Codeid: Input.Codeid,
	}).First(&admin)

	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"errpr": "User not ecits"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(admin.Pass), []byte(Input.Pass))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Login ": "Fail"})
		return
	}

	k := JwtToken.NewJWTService().GenerateToken(Input.Codeid, admin.Codeid)
	c.JSON(http.StatusOK, gin.H{"Login successful": k})
	return
}

/////
func RegisterUser(c *gin.Context) {
	var user InforJwtB3.Staffs
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error ": err.Error()})
		c.Abort()
		return
	}
	if err := user.HashPassword(user.Pass); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error ": err.Error()})
		c.Abort()
		return
	}
	record := ConnectJwtB3.DB.Create(&user)

	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error ": record.Error.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Codeid": user.Codeid, "Pass": user.Pass})
}

////auth
type TokenResquet struct {
	Codeid string `json:"codeid"`
	Pass   string `json:"pass"`
}

func GenerateToken(c *gin.Context) {
	var request TokenResquet
	var user InforJwtB3.Staffs
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		c.Abort()
		return
	}
	record := ConnectJwtB3.DB.Where("Codeid = ?", request.Codeid).Find(&user)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error ": record.Error.Error()})
		c.Abort()
		return
	}
	credentislError := user.CheckPassword(request.Pass)
	if credentislError != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error ": "invalid credentials"})
		c.Abort()
		return
	}
	tokenString, err := TokenJwtB3.GenerateJWTauth(user.Codeid, user.Pass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error ": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"token ": tokenString})
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message ": "Pong"})
}
