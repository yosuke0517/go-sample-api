FROM golang:1.13.0-alpine3.10 as build

ENV GOPATH /go

WORKDIR $GOPATH/src

ENV GO111MODULE=on

WORKDIR $GOPATH/src/app

COPY . .

RUN set -ex && \
    apk update && \
    apk add --no-cache git && \
    go build -o portfolio-backend && \
    go get gopkg.in/urfave/cli.v2@master && \
    go get github.com/oxequa/realize && \
    go get -u github.com/go-delve/delve/cmd/dlv && \
    go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv

FROM alpine:3.10

COPY --from=build /go/app/app .

RUN set -x \
    && addgroup go \
    && adduser -D -G go go \
    && chown -R go:go /app/app

CMD ["./app"]
