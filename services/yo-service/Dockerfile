FROM golang:1.12.5-alpine
MAINTAINER James Turnbull <james@lovedthanlost>

ENV service_path /go/src/github.com/jamtur01/yoservice/
ENV PATH $PATH:$service_path

WORKDIR $service_path
COPY yoservice/ .

RUN go build .

ENTRYPOINT ./yoservice

EXPOSE 8080
