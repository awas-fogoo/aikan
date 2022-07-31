FROM golang:latest

MAINTAINER aikan_admin
WORKDIR /awesomeProject0511
WORKDIR $GOPATH/go
COPY . $GOPATH/go
COPY go.mod ./
COPY go.sum ./
COPY * ./
RUN go env -w GO111MODULE=auto
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -mod=mod main.go routes.go


EXPOSE 8888
CMD ["./main"]
