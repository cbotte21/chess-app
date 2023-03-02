# syntax=docker/dockerfile:1

FROM golang:latest
MAINTAINER cbotte21@gmail.com

WORKDIR /app

COPY * ./

RUN go build -o server
CMD ["./server"]
EXPOSE 5000
