FROM golang:1.19-alpine AS builder

COPY . /go/src/github.com/tigercandy/prado

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build .

FROM alpine:latest

WORKDIR /go/src/github.com/tigercandy/prado

RUN echo -e https://mirrors.ustc.edu.cn/alpine/v3.13/main/ > /etc/apk/repositories

RUN apk --no-cache add tzdata &&\
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

COPY --from=builder /opt/repo/src/github.com/tigercandy/prado /opt

WORKDIR /opt

EXPOSE 8088
ENTRYPOINT ./server -c configs/config.yaml