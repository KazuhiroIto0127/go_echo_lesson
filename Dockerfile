# -----------
# build環境
# -----------
FROM golang:1.21.1-alpine as builder

# ログに出力する時間をJSTにするため、タイムゾーンを設定
ENV TZ /usr/share/zoneinfo/Asia/Tokyo

WORKDIR /app
COPY ./ /app

RUN apk add --no-cache build-base make
RUN apk add --no-cache git sqlite-dev gcc musl-dev

RUN go install github.com/cosmtrek/air@v1.27.3
RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install golang.org/x/tools/gopls@latest

RUN go build -o main ./cmd/myapp/main.go

EXPOSE 8080

# -----------
# production環境
# -----------
FROM alpine as production
RUN apk update && apk upgrade
RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/main ./main
COPY --from=builder /app/example.sql ./example.sql

CMD ["./main"]
