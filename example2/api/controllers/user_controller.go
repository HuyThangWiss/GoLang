package controllers

import (
	"example2/api/request"
	"example2/core/entities"
	"example2/core/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}
func (u *UserController) CreateUser(c *gin.Context) {
	var createUserReq request.Books
	if err := c.ShouldBindJSON(&createUserReq); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, "Loi roi")
		return
	}
	_, err := u.userService.CreateUser(c, &createUserReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "loi")
		return
	}
	c.JSON(http.StatusOK, createUserReq)
}

func (u *UserController) FinAll(c *gin.Context) {
	arr := make([]entities.Books, 0)
	arr, err := u.userService.FindAllUser(c, arr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "loi")
		return
	}
	c.JSON(http.StatusOK, gin.H{"Data ": arr})
}

func (u *UserController) Fin_Id(c *gin.Context) {
	Id := c.Param("Id")
	arr := make([]entities.Books, 0)
	arr, err := u.userService.Find_Id(c, Id, arr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "err")
		return
	}
	if len(arr) ==0{
		c.JSON(http.StatusInternalServerError, "Data not exists")
		return
	}
	c.JSON(http.StatusOK, gin.H{"Data ": arr})
}
func (u *UserController) UpDateUser(c *gin.Context) {
	Id := c.Param("Id")
	var UpdateUserReq request.Books
	arr := make([]entities.Books, 0)

	if err := c.ShouldBindJSON(&UpdateUserReq); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, "Loi roi")
		return
	}

	arr, err := u.userService.Find_Id(c, Id, arr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "err")
		return
	}
	if len(arr) ==0{
		c.JSON(http.StatusInternalServerError, "Data not exists")
		return
	}
	err1 := u.userService.Update(c, Id, &request.Books{
		Title:       UpdateUserReq.Title,
		Description: UpdateUserReq.Description,
	})
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}
	c.JSON(http.StatusOK, gin.H{"After update ": UpdateUserReq})

}

func (u *UserController) DeleteUser(c *gin.Context) {
	Id := c.Param("Id")

	arr := make([]entities.Books, 0)
	arr, err := u.userService.Find_Id(c, Id, arr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}
	if len(arr) ==0{
		c.JSON(http.StatusInternalServerError, "Data not exists")
		return
	}
	err1 := u.userService.Delete(c, nil, Id)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, "Not delete")
		return
	}
	c.JSON(http.StatusLocked, "Delete successfully")
}
