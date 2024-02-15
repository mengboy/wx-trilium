FROM golang:alpine AS builder

ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY . /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o wx-trilium .

FROM alpine:latest
LABEL org.opencontainers.image.source="https://github.com/mengboy/wx-trilium"

COPY --from=builder /app/wx-trilium .

ENV GIN_MODE=release
EXPOSE 1234
VOLUME /conf.toml

ENTRYPOINT ["/wx-trilium", "start"]
