# go-sample-api

- 最初だけ
  - ~~`docker run -v`pwd`:/go/src/app -w /go/src/app golang:1.13.0-alpine go mod init app`~~
  - ~~`docker build -t myapp .`~~
  - ~~`docker run -p 8080:8080 -d --name app app`~~
  - 現状だと`docker-compose build`・`docker-compose up`で動く。ハズ・・