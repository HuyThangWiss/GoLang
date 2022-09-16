package main

import (
	"b1/JWT/B1/ConnectJwtB1"
	"b1/JWT/B1/ControllersJwtB1"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	ConnectJwtB1.ConnectJwtB1()
	r.GET("/Select", ControllersJwtB1.SelectJwtB1)
	r.POST("/post", ControllersJwtB1.GenerateToken)
	r.POST("/register", ControllersJwtB1.RegisterUser)
	r.POST("/Insert", ControllersJwtB1.InsertJwtB1)
	r.Run()

}
