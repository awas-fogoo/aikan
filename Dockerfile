FROM golang:1.17-alpine

MAINTAINER foo
WORKDIR /AIKAN
COPY . /AIKAN
COPY go.mod ./
COPY go.sum ./

RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod tidy \
    && go build -mod=mod main.go \
    && sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk update --no-cache \
    && apk add ffmpeg


EXPOSE 12333
CMD ["./main"]