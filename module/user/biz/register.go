package biz

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/module/user/model"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*model.User, error)
	CreateUser(ctx context.Context, data *model.UserCreate) error
}

type Hasher interface {
	Hash(password string, salt []byte) string
	ValidatePassword(hashPassword, currentPassword string, salt []byte) bool
}

type registerBusiness struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterBusiness(registerStorage RegisterStorage, hasher Hasher) *registerBusiness {
	return &registerBusiness{
		registerStorage: registerStorage,
		hasher:          hasher,
	}
}

func (business *registerBusiness) Register(ctx context.Context, data *model.UserCreate) error {
	user, _ := business.registerStorage.FindUser(ctx, map[string]interface{}{"username": data.Username})
	if user != nil {
		return model.ErrUsernameExisted
	}

	salt := common.GenSalt(50)

	data.Password = business.hasher.Hash(data.Password, salt)
	data.Salt = string(salt)
	// data.Role = "user" harod
	default_role := "user"
	if data.Role == nil {
		data.Role = &default_role
	}
	if err := business.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}
	return nil
}
