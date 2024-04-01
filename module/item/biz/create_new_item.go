package biz

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/module/item/model"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoItemCreation) error
}

type createItemBiz struct {
	storage CreateItemStorage
}

func NewCreateItemBiz(store CreateItemStorage) *createItemBiz {
	return &createItemBiz{storage: store}
}

func (biz *createItemBiz) CreateNewItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.storage.CreateItem(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}
	return nil
}
