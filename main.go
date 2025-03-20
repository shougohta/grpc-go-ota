package main

import (
	handler "example.com/mod/internal/handler/grpc"
	grpc "example.com/mod/internal/infrastructure/grpc"
	persistence "example.com/mod/internal/infrastructure/persistence"
	service "example.com/mod/internal/service"
)

func main() {
	userRepo := persistence.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	grpc.StartGRPCServer(userHandler)
}
