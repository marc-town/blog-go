FROM golang:latest AS build

WORKDIR /go/src/github.com/marc-town/blog-go/web-api/app
ENV GO111MODULE=on
COPY ./app/go.mod .
COPY ./app/go.sum .
RUN go mod download
COPY ./app .

RUN go build -o /go/bin/app \
    && go get gopkg.in/urfave/cli.v2@master \
    && go get github.com/oxequa/realize
