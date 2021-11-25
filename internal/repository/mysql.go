package repository

import (
	"context"

	"github.com/kianooshaz/cleanservice/internal/entity/model"
)

type Mysql interface {
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	DeleteUser(ctx context.Context, id int) error
}
