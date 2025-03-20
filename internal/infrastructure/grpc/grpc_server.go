package grpc

import (
	"log"
	"net"

	handler "example.com/mod/internal/handler/grpc"
	proto "example.com/mod/internal/proto"
	"google.golang.org/grpc"
)

// StartGRPCServer は gRPC サーバーを起動する（Infrastructure Layer）
func StartGRPCServer(userHandler *handler.UserHandler) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterUserServiceServer(grpcServer, userHandler)

	log.Println("gRPC server running on port 50051")
	grpcServer.Serve(lis)
}
