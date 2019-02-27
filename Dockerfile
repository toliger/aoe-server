FROM golang:alpine AS build-env

ENV GO111MODULE=on

RUN apk add git

WORKDIR /go/src/app

COPY . .

RUN go build


FROM alpine:latest

WORKDIR /app

COPY --from=build-env /go/src/app/server  /app/server

RUN chmod +x /app/server

ENTRYPOINT ["./server"]
