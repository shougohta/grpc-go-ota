package domain

import (
	"context"

	model "example.com/mod/internal/domain/model"
)

// UserRepository はインフラ層からユーザーデータの取得を行うためのインターフェース
type UserRepository interface {
	FindByID(ctx context.Context, id string) (*model.User, error)
}
