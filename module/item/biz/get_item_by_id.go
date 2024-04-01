package biz

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/module/item/model"
)

type GetItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
}

type getitemBiz struct {
	store GetItemStorage
}

func NewGetItemBiz(store GetItemStorage) *getitemBiz {
	return &getitemBiz{store: store}
}

func (biz *getitemBiz) GetItemById(ctx context.Context, id int) (*model.TodoItem, error) {
	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, common.ErrCannotGetEntity(model.EntityName, err)
	}
	return data, nil
}
