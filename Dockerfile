FROM golang:1.13-alpine

# 作業ディレクトリの作成・設定
WORKDIR /panenga-api
# Go Modules を有効化
ENV GO111MODULE=on

COPY go.mod .

# go.mod 内のパッケージをダウンロード
RUN go mod download

EXPOSE 8080