FROM golang:1.12-alpine AS build-env

ENV GO111MODULE=on

RUN apk add --no-cache git

WORKDIR ${GOPATH}/src/github.com/container-examples/golang-webserver/
COPY . .

RUN go mod download
RUN go build -ldflags="-w -s" -o ./build/webserver .

FROM alpine:3.11

WORKDIR /app
COPY --from=build-env /go/src/github.com/container-examples/golang-webserver/build/webserver /app/webserver

ENTRYPOINT [ "/app/webserver" ]