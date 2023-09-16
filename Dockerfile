FROM golang:1.21.1

# ログに出力する時間をJSTにするため、タイムゾーンを設定
ENV TZ /usr/share/zoneinfo/Asia/Tokyo

WORKDIR /app
COPY ./ /app

RUN apt-get update && apt-get install -y build-essential make

RUN go install github.com/cosmtrek/air@v1.27.3
RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install golang.org/x/tools/gopls@latest

RUN go build

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

EXPOSE 8080

CMD ["air"]
