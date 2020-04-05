FROM golang:1.13 as build

ENV GOPATH /go

WORKDIR $GOPATH/src

ENV GO111MODULE=on

WORKDIR $GOPATH/src/app

COPY . .

RUN set -ex && \
  apt-get update && \
  apt-get install -y git \
  vim && \
  go build -o video-manager-go && \
  # realizeはメンテされてないからfreshを使用
  # go get github.com/urfave/cli && \
  # go get github.com/urfave/cli/v2 && \
  # go get -u github.com/oxequa/realize && \
  # go get github.com/oxequa/realize && \
  go get github.com/pilu/fresh && \
  go get -u github.com/go-delve/delve/cmd/dlv && \
  go get github.com/PuerkitoBio/goquery && \
  go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv

FROM alpine:3.10

COPY --from=build go/src/app/video-manager-go .

RUN set -x \
    && addgroup go \
    && adduser -D -G go go \
    && chown -R go:go src/app/video-manager-go

CMD ["./video-manager-go"]
# CMD ["fresh"]