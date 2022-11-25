package port

import (
	"context"
	"example2/adapters/databases"
	"example2/core/entities"
)

type UserRepositoryPort interface {
	InsertUser(ctx context.Context, user *entities.Books) error
	FindAllUser(ctx context.Context, user []entities.Books) ([]entities.Books, error)
	FindUser(ctx context.Context, user []entities.Books, id string) ([]entities.Books, error)
	UpdateUser(ctx context.Context,id string, user *entities.Books) error
	DeleteUser(ctx context.Context, user *entities.Books,Id string) error
}

func InitUserRepositoryPort(Db *databases.PostgreSql) UserRepositoryPort {
	return Db
}
