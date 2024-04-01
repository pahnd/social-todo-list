package storage

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/module/item/model"
)

func (s *sqlStore) GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
