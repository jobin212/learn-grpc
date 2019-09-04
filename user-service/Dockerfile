# consignment-service/Dockerfile
FROM alpine:latest

RUN mkdir -p /go/src/github.com/jobin212/shippy

WORKDIR /go/src/github.com/jobin212/shippy

ADD user-service /go/src/github.com/jobin212/shippy/user-service

CMD ["./user-service"]