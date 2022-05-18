FROM golang:1.18.2-alpine as migrate
RUN apk update \
    && apk add --no-cache --virtual goose-dependencies \
         build-base \
         git \
    && GO111MODULE=on go get bitbucket.org/liamstask/goose/cmd/goose \
    && apk del goose-dependencies

FROM golang:1.18.2-alpine as builder
WORKDIR /app
ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN apk update \
  && apk --no-cache add git tzdata
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build \
    -o /app/build \
    -ldflags '-s -w'

FROM alpine:3.15.0
WORKDIR /app
COPY --from=migrate /go/bin/goose /usr/local/bin/
COPY --from=builder /usr/share/zoneinfo/Asia/Tokyo /usr/share/zoneinfo/Asia/Tokyo
COPY --from=builder /app/build /app
COPY ./entrypoint.sh /app
COPY ./db /app/db
COPY ./src/config/ /app/src/config/
RUN apk update \
    && apk add --no-cache mariadb-client \
    && rm /var/cache/apk/*
ENTRYPOINT ["sh", "./entrypoint.sh"]
