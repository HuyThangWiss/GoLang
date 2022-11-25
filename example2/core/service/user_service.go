package service

import (
	"context"
	"example2/api/request"
	"example2/core/entities"
	"example2/core/port"
)

type UserService struct {
	userRepositoryPort port.UserRepositoryPort
}

func NewUserService(userRepositoryPort port.UserRepositoryPort) *UserService {
	return &UserService{
		userRepositoryPort: userRepositoryPort,
	}
}

func (u *UserService) CreateUser(ctx context.Context, req *request.Books) (*entities.Books, error) {
	err := u.userRepositoryPort.InsertUser(ctx, &entities.Books{
		Id:          req.Id,
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (u *UserService) FindAllUser(ctx context.Context, arr []entities.Books) ([]entities.Books, error) {

	arr, err := u.userRepositoryPort.FindAllUser(ctx, arr)

	if err != nil {
		return nil, nil
	}
	return arr, nil
}

func (u *UserService) Find_Id(ctx context.Context, Id string, arr []entities.Books) ([]entities.Books, error) {

	arr, err := u.userRepositoryPort.FindUser(ctx, arr, Id)
	if err != nil {
		return nil, nil
	}
	return arr, nil
}

func (u *UserService) Update(ctx context.Context, id string, user *request.Books) error {
	err := u.userRepositoryPort.UpdateUser(ctx, id, &entities.Books{
		Id:          user.Id,
		Title:       user.Title,
		Description: user.Description,
	})
	if err != nil {
		return nil
	}
	return nil
}

func (u *UserService) Delete(ctx context.Context, user *entities.Books, Id string) error {
	err := u.userRepositoryPort.DeleteUser(ctx, user, Id)
	if err != nil {
		return nil
	}
	return nil
}
