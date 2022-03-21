FROM golang:1.17-alpine as builder

WORKDIR /app

RUN apk update && \
    apk add gcc musl-dev upx
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o eldencli cmd/eldencli/main.go
RUN upx -q -9 eldencli


FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/eldencli .
ENTRYPOINT [ "./eldencli", "serve" ]