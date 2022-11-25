package databases

import (
	"context"
	"example2/adapters/databases/model"
	"example2/core/entities"
	"fmt"
	"gorm.io/gorm"
)

//type DB struct {
//	Db *gorm.DB
//}

type PostgreSql struct {
	DB *gorm.DB
}

func NewPostGre(DB *gorm.DB) *PostgreSql {
	return &PostgreSql{
		DB: DB,
	}
}

func (postGr *PostgreSql) InsertUser(ctx context.Context, user *entities.Books) error {

	userModel := &model.Books{
		Id:          user.Id,
		Title:       user.Title,
		Description: user.Description,
	}
	resul := postGr.DB.Create(&userModel)
	if resul.Error != nil {
		fmt.Println("Hello")
		return nil
	}
	return nil
}

func (postGr *PostgreSql) FindAllUser(ctx context.Context, user []entities.Books) ([]entities.Books, error) {

	result := postGr.DB.Find(&user)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return nil, nil
	}
	return user, nil
}

func (postGr *PostgreSql) FindUser(ctx context.Context, user []entities.Books, id string) ([]entities.Books, error) {

	result := postGr.DB.Where("Id = ? ", id).Find(&user)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return nil, nil
	}
	return user, nil
}

func (postGr *PostgreSql) UpdateUser(ctx context.Context, id string, user *entities.Books) error {
	userModel := model.Books{
		Id:          user.Id,
		Title:       user.Title,
		Description: user.Description,
	}
	resul := postGr.DB.Where("Id = ?", id).Updates(&userModel)
	if resul.Error != nil {
		fmt.Println("Err")
		return nil
	}

	return nil
}
func (postGr *PostgreSql) DeleteUser(ctx context.Context, user *entities.Books, Id string) error {

	resul := postGr.DB.Where("Id = ? ", Id).Delete(&user)
	if resul.Error != nil {
		fmt.Println("Err")
		return nil
	}
	return nil
}
