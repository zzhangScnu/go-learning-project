package service

import (
	"context"
	"github.com/go-learning-project/gen/config"
	"github.com/go-learning-project/gen/dal/model"
	"github.com/go-learning-project/gen/dal/query"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u UserService) QueryAll(ctx context.Context) ([]*model.TUser, error) {
	user := query.Use(config.DB).TUser
	return user.WithContext(ctx).Find()
}

func (u UserService) QueryByName(name string, ctx context.Context) (map[string]interface{}, error) {
	user := query.Use(config.DB).TUser
	return user.WithContext(ctx).FindByName(name)
}
