# consignment-service/Dockerfile
FROM alpine:latest

RUN mkdir -p /go/src/github.com/jobin212/shippy

WORKDIR /go/src/github.com/jobin212/shippy

ADD vessel-service /go/src/github.com/jobin212/shippy/vessel-service

CMD ["./vessel-service"]