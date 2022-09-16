package main

import (
	"b1/JWT/B3/ConnectJwtB3"
	"b1/JWT/B3/ControllerJwtB3"
	"b1/JWT/B3/MiddwareJwtB3"
	"github.com/gin-gonic/gin"
)

func main() {
	ConnectJwtB3.ConnectJwtB3()
	r := gin.Default()
	//
	//api :=r.Group("/api",MiddwareJwtB3.AuthorizeJWT())
	//{
	//	api.GET("/Select",ControllerJwtB3.Select)
	//	api.POST("/Insert",ControllerJwtB3.Insert)
	//}
	//
	//r.POST("/Login",ControllerJwtB3.Logint)
	api := r.Group("/api")
	{
		api.GET("/Select", ControllerJwtB3.Select)
		api.POST("/token", ControllerJwtB3.GenerateToken)
		api.POST("/user", ControllerJwtB3.RegisterUser)
		secured := api.Group("/secured").Use(MiddwareJwtB3.Auth())
		{
			secured.GET("/ping", ControllerJwtB3.Ping)
		}
	}
	r.Run()
}
