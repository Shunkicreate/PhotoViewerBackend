# ビルドステージ
FROM golang:1.21-alpine AS builder

# 作業ディレクトリを設定
WORKDIR /app

# ソースコードをコピー
COPY . .

# 必要なモジュールをインストール
RUN go mod tidy

# バイナリをビルド（デバッグ情報を除去）
RUN GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o api-server .

# 実行ステージ
FROM gcr.io/distroless/base

# バイナリをコピー
COPY --from=builder /app/api-server /api-server

# 実行時にコマンドを指定
CMD ["/api-server"]
