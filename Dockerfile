FROM golang:alpine AS build-env

ENV GO111MODULE=on

RUN apk add git

WORKDIR /go/src/app

COPY . .

RUN go build -o /bin/server


FROM scratch

COPY --from=build-env /bin/server /bin/server

ENTRYPOINT ["/bin/server"]

