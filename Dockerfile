FROM golang:alpine AS builder
MAINTAINER kyobum.keum <kyobum.keum@gmail.com>

RUN mkdir -p /go/src/github.com/gin-practice
ADD . /go/src/github.com/gin-practice

WORKDIR /go/src/github.com/gin-practice

RUN go mod tidy
RUN go build main.go

### Executable Image
FROM alpine

COPY --from=builder /go/src/github.com/gin-practice/main ./main
EXPOSE 8080

ENTRYPOINT ["./main"]