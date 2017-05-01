FROM golang:1.8.1-alpine

ENV webserver_path /go/src/github.com/perriea/webserver/
ENV PATH $PATH:$webserver_path

WORKDIR $webserver_path
ADD webserver/ .
RUN go build . && \
    mv webserver /go/bin/

ENTRYPOINT /go/bin/webserver

EXPOSE 8080
