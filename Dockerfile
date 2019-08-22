FROM golang:1.12-alpine AS build-env

LABEL ROLE "build"
ENV GO111MODULE=off

RUN apk add --no-cache make

WORKDIR ${GOPATH}/src/github.com/container-examples/golang-webserver/
COPY . .

RUN make

FROM alpine:3.10
LABEL maintainer="a.perrier89@gmail.com"

COPY --from=build-env /go/src/github.com/container-examples/golang-webserver/build/webserver /bin/webserver

ENTRYPOINT [ "/bin/webserver" ]