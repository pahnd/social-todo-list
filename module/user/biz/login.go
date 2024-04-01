package biz

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/component/tokenprovider"
	"social-todo-list/module/user/model"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*model.User, error)
}

type loginBussiness struct {
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(storeUser LoginStorage, tokenProvider tokenprovider.Provider,
	hasher Hasher, expiry int) *loginBussiness {
	return &loginBussiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

func (business *loginBussiness) Login(ctx context.Context, data *model.UserLogin) (tokenprovider.Token, error) {
	user, err := business.storeUser.FindUser(ctx, map[string]interface{}{"username": data.Username})

	if err != nil {
		return nil, model.ErrUsernameOrPasswordInvalid
	}

	if business.hasher.ValidatePassword(user.Password, data.Password, []byte(user.Salt)) != true {
		return nil, model.ErrUsernameOrPasswordInvalid
	}

	payload := &common.TokenPayload{
		UId:   user.Id,
		URole: user.Role.String(),
	}

	accessToken, err := business.tokenProvider.Generate(payload, business.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return accessToken, nil
}
