package persistence

import (
	"context"
	"errors"

	model "example.com/mod/internal/domain/model"
	repository "example.com/mod/internal/domain/repository"
)

// UserRepositoryImpl は UserRepository の実装
type UserRepositoryImpl struct {
	id string
}

// NewUserRepository は UserRepositoryImpl を生成
func NewUserRepository() repository.UserRepository {
	return &UserRepositoryImpl{id: "123"}
}

// FindByID はDBからユーザーを取得（仮実装）
func (r *UserRepositoryImpl) FindByID(ctx context.Context, id string) (*model.User, error) {
	if id == r.id {
		return &model.User{ID: r.id, Name: "Taro"}, nil
	}
	return nil, errors.New("user not found")
}
