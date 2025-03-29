# gRPCサーバー for オニオンアーキテクチャ

## ディレクトリ構造

```
myapp/
├── README.md
├── client // クライアントサーバー
│   └── main.go
├── cmd
├── go.mod
├── go.sum
├── internal
│   ├── domain  // ドメインサービス層
│   │   ├── model
│   │   │   └── user.go
│   │   └── repository // インフラ層へのアクセス(依存性逆転)
│   │       └── user_repository.go
│   ├── grpc
│   ├── handler // ハンドラー層
│   │   └── grpc
│   │       └── user_handler.go
│   ├── infrastructure // インフラ層
│   │   ├── grpc
│   │   │   └── grpc_server.go
│   │   └── persistence
│   │       └── user_repository.go
│   ├── proto // プロトコルバッファ
│   │   ├── user.pb.go
│   │   ├── user.proto
│   │   └── user_grpc.pb.go
│   └── service
│       └── user_service.go
└── main.go
```

## クライアントコード生成

```
protoc --go_out=. --go-grpc_out=. internal/proto/user.proto
```

## gRPCサーバー起動

rootディレクトリで以下を実行。

```
go run main.go
```

## クライアントサーバー起動

```
cd client
```

```
go run main.go
```