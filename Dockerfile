FROM golang:1.20-alpine as builder
WORKDIR /source
EXPOSE 8888
COPY . .
RUN go build .


