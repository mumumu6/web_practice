# 使用するベースイメージ
FROM golang:1.20 AS builder

# 作業ディレクトリの設定
WORKDIR /app

# Go モジュールのコピーと依存関係のインストール
COPY go.mod go.sum ./
RUN go mod download

# ソースコードのコピー
COPY . .

# アプリケーションのビルド
RUN go build -o main .

# 実行用のベースイメージ
FROM alpine:latest

# 作業ディレクトリの設定
WORKDIR /root/

# ビルドしたバイナリのコピー
COPY --from=builder /app/main .

# アプリケーションの実行
CMD ["./main"]
