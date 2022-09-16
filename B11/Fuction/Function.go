package Fuction

import (
	"b1/B11/Connect"
	"b1/B11/Infor"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Select(c *gin.Context) {
	Connect.Connect()
	var arr []Infor.Admins
	Connect.DB.Find(&arr)
	c.JSON(http.StatusOK, gin.H{"Data ": arr})

}

func Insert(c *gin.Context) {
	Connect.Connect()
	var Input Infor.Admins
	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error ": err.Error()})
		return
	}

	human := Infor.Admins{Useradmin: Input.Useradmin, Passwordadmin: Input.Passwordadmin}
	Connect.DB.Create(&human)
	c.JSON(http.StatusOK, gin.H{"Data : ": human})
}

func EnCode(password string) (string, error) {

	var passwordBytes = []byte(password)

	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	return string(hashedPasswordBytes), err
}

func SignIn(c *gin.Context) {
	Connect.Connect()
	var Input Infor.Admins

	//	var arr[] Infor.Admins
	//	Connect.DB.Find(&arr)

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error : ": err.Error()})
		return
	}

	var admin Infor.Admins
	//	hashPassword, _ := EnCode(Input.Passwordadmin)
	if Input.Useradmin == "" || Input.Passwordadmin == "" {
		c.JSON(http.StatusOK, gin.H{"Login ": "Input Empty"})
		return
	}
	/*
		else if err2 := Connect.DB.Where("Useradmin = ?", Input.Useradmin).First(&a).Error; err2 == nil {

			if err3 := Connect.DB.Where("Passwordadmin = ?", hashPassword).First(&a).Error; err3 == nil {
				c.JSON(http.StatusOK, gin.H{"Login ": "Successful"})
			} else {
				c.JSON(http.StatusOK, gin.H{"Login ": "Err pass"})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"Login ": "User not exits"})
		}
	*/
	////////////////
	result := Connect.DB.Where(&Infor.Admins{
		Useradmin: Input.Useradmin,
	}).First(&admin)
	if result.Error != nil {
		// handle err
		return
	}
	//compare password req và pasword của admin (db)
	err := bcrypt.CompareHashAndPassword([]byte(admin.Passwordadmin), []byte(Input.Passwordadmin))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Login ": "Fail"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Login ": "Successful"})

}

func DecCode(hashedPassword, currPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(currPassword))
	return err == nil
}

func DecLogin(c *gin.Context) {
	Connect.Connect()
	var Input Infor.Admins

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error : ": err.Error()})
		return
	}
	var admin Infor.Admins
	result := Connect.DB.Where(&Infor.Admins{
		Useradmin: Input.Useradmin,
	}).First(&admin)
	if result.Error != nil {
		// handle err
		return
	}
	//compare password req và pasword của admin (db)
	err := bcrypt.CompareHashAndPassword([]byte(admin.Passwordadmin), []byte(Input.Passwordadmin))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Login ": "Fail"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Login ": "Successful"})

}

func Hasjpassword(c *gin.Context) {
	Connect.Connect()
	var Input Infor.Admins

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error ": err.Error()})
		return
	}
	hashedPassword, err := EnCode(Input.Passwordadmin)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Error": err.Error()})
	}
	Human := Infor.Admins{Useradmin: Input.Useradmin, Passwordadmin: hashedPassword}
	Connect.DB.Create(&Human)

	c.JSON(http.StatusOK, gin.H{"Data ": Human})

}
