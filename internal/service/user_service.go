package service

import (
	"context"

	model "example.com/mod/internal/domain/model"
	repository "example.com/mod/internal/domain/repository"
)

// UserService はユーザー情報を扱うユースケース
type UserService struct {
	userRepo repository.UserRepository
}

// NewUserService は UserService を生成
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

// GetUser はユーザー情報を取得する
// repositoryをインプットと、アウトプットとしてモデルを返す
func (s *UserService) GetUser(ctx context.Context, id string) (*model.User, error) {
	return s.userRepo.FindByID(ctx, id)
}
