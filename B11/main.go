package main

import (
	"b1/B11/Fuction"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/Select", Fuction.Select)
	r.POST("/Insert", Fuction.Insert)
	r.POST("/InsertEnpass", Fuction.Hasjpassword)
	r.POST("/Login", Fuction.SignIn)
	r.POST("/LoginDec", Fuction.DecLogin)
	r.Run()

}
