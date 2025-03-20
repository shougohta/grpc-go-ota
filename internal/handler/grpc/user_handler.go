package grpc

import (
	"context"

	proto "example.com/mod/internal/proto"
	service "example.com/mod/internal/service"
)

// UserHandler は gRPC のエントリーポイント（Handlerレイヤー）
type UserHandler struct {
	proto.UnimplementedUserServiceServer
	userService *service.UserService
}

// NewUserHandler は UserHandler を生成
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// GetUser は gRPC リクエストを処理する
func (h *UserHandler) GetUser(ctx context.Context, req *proto.UserRequest) (*proto.UserResponse, error) {
	user, err := h.userService.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &proto.UserResponse{Id: user.ID, Name: user.Name}, nil
}
