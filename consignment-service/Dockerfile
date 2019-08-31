# consignment-service/Dockerfile
FROM alpine:latest

RUN mkdir -p /go/src/github.com/jobin212/shippy

WORKDIR /go/src/github.com/jobin212/shippy

ADD consignment-service /go/src/github.com/jobin212/shippy/consignment-service

CMD ["./consignment-service"]