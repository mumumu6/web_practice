# ベースイメージとして Golang を使用
FROM golang:1.22

# 作業ディレクトリを設定
WORKDIR /go/app

# 必要なパッケージをインストール
RUN go install github.com/air-verse/air@latest

# コンテナが起動したときに実行するコマンドを指定
CMD ["air", "-c", ".air.toml"]
