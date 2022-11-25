package main

import (
	"example2/adapters/databases"
	"example2/api/controllers"
	"example2/core/port"
	"example2/core/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)


func main() {

	postGresCollect := NewPostGresCollect()
	postG := databases.NewPostGre(postGresCollect)
	userRepositoryPort := port.InitUserRepositoryPort(postG)
	userService := service.NewUserService(userRepositoryPort)
	userController := controllers.NewUserController(userService)

	r := gin.Default()
	r.POST("/user",userController.CreateUser)
	r.GET("/get",userController.FinAll)
	r.GET("/Search/:Id",userController.Fin_Id)
	r.PUT("/update/:Id",userController.UpDateUser)
	r.DELETE("/delete/:Id",userController.DeleteUser)
	if err := r.Run();err != nil{
		fmt.Println(err)
		return
	}
}

func NewPostGresCollect()*gorm.DB  {

	dsn := "host=localhost user=postgres password=1234 dbname=Books port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connect fail")
	} else {
		fmt.Print("Connect successfully")
	}

	return db
}