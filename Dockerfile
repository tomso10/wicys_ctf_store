FROM golang:1.19.3-alpine

WORKDIR /app
COPY . .

RUN go get ./...
RUN go mod tidy
RUN apk add build-base

RUN go build -o build/main main.go

CMD ["./build/main"]
