package mysql

import (
	"context"

	"github.com/kianooshaz/cleanservice/internal/entity/model"
)

func (m *mysql) CreateUser(ctx context.Context, user *model.User) error {
	return nil
}

func (m *mysql) UpdateUser(ctx context.Context, user *model.User) error {
	return nil
}

func (m *mysql) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	return nil, nil
}

func (m *mysql) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	return nil, nil
}

func (m *mysql) DeleteUser(ctx context.Context, id int) error {
	return nil
}
