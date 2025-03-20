# オニオンアーキテクチャのgRPCサーバー

## ディレクトリ構造

myapp/
│── cmd/               # アプリのエントリーポイント（main.go）
│── internal/
│   ├── domain/        # ドメインレイヤー（Entity, Repositoryのインターフェース）
│   ├── service/       # ビジネスロジック（ユースケース）
│   ├── repository/    # インフラ層（DBリポジトリ）
│   ├── grpc/          # gRPCハンドラー
│   ├── proto/         # Protocol Buffers定義と生成コード
│── infrastructure/    # 外部依存（DB, gRPCサーバー）
│── client/            # gRPCクライアント（クライアントアプリケーション）
│   └── main.go        # クライアント実行ファイル
│── go.mod             # Go モジュール管理
│── main.go            # gRPCサーバーのエントリーポイント

## クライアントコード生成

```
protoc --go_out=. --go-grpc_out=. internal/proto/user.proto
```