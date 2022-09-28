package ControllerM

import (
	"awesomeProject/MonoBb/MonoB11/ConnectMoNoB11"
	"awesomeProject/MonoBb/MonoB11/InforM"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Firstname string `json:"firstname"`
//Lastname string `json:"lastname"`
//Gender string `json:"gender"`
//Age int `json:"age"`
//Address string `json:"address"`
//Gmail string `json:"gmail"`

type Req struct {
	Firstname string `form:"firstname"`
	Lastname  string `form:"lastname"`
	Gender    string `form:"gender"`
	Age       int    `form:"age"`
	Address   string `form:"address"`
	Gmail     string `form:"gmail"`
}

func Select(c *gin.Context) {

	ConnectMoNoB11.ConnectM()
	var arr []InforM.Users
	var req Req
	err := c.ShouldBindQuery(&req)
	if err != nil {
		fmt.Println(err)
		return
	}

	//ConnectMoNoB11.DB.Where("firstname = ? and lastname = ? and gender = ? and age = ? and address = ? and gmail = ? ",
	//	req., lastname, gender, age, address, gmail).Find(&arr)
		result := ConnectMoNoB11.DB.Where(InforM.Users{
		Firstname: req.Firstname,
		Lastname: req.Lastname,
		Gender:    req.Gender,
		Age:       req.Age,
		Address:   req.Address,
		Gmail:     req.Address,
	}).Find(&arr)
	if result.Error!= nil {
		fmt.Println(result.Error.Error())
		return
	}
	c.JSON(http.StatusOK,gin.H{"data ":arr})
}
