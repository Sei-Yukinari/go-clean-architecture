FROM golang:1.18.2-alpine as migrate
RUN apk update \
    && apk add --no-cache --virtual goose-dependencies \
         build-base \
    && go install bitbucket.org/liamstask/goose/cmd/goose@latest \
    && apk del goose-dependencies

FROM golang:1.18.2-alpine as builder
WORKDIR /app
# rdb migtation
COPY --from=migrate /go/bin/goose /usr/local/bin/goose
COPY ./db /app/db
RUN goose up
# app build
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

FROM alpine:3.15
WORKDIR /app
COPY --from=builder /usr/share/zoneinfo/Asia/Tokyo /usr/share/zoneinfo/Asia/Tokyo
COPY --from=builder /app/build /app
COPY ./src/config/ /app/src/config/
ENTRYPOINT ["/app/build"]
