FROM golang:alpine AS build-env

ENV GO111MODULE=on

RUN apk add git

WORKDIR /go/src/app

COPY . .

RUN GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -o /go/bin/server


FROM scratch

COPY --from=build-env /go/bin/server /go/bin/

COPY data/SmallTestMap.json /go/bin/data/SmallTestMap.json

WORKDIR /go/bin

EXPOSE 50010/tcp

ENTRYPOINT ["/go/bin/server"]

