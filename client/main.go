package main

import (
	"context"
	"log"
	"time"

	proto "example.com/mod/internal/proto"

	"google.golang.org/grpc"
)

func main() {
	// gRPCサーバーに接続
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// gRPCクライアントの作成
	client := proto.NewUserServiceClient(conn)

	// コンテキストを作成（タイムアウトを設定）
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// gRPC API を呼び出し
	res, err := client.GetUser(ctx, &proto.UserRequest{Id: "123"})
	if err != nil {
		log.Fatalf("could not fetch user: %v", err)
	}

	// レスポンスを表示
	log.Printf("User: ID=%s, Name=%s", res.Id, res.Name)
}
