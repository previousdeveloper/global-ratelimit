FROM registry.trendyol.com/platform/base/image/golang:1.13.4-alpine3.10 AS builder

ENV GOPATH /go
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ARG VERSION

RUN mkdir /app
WORKDIR /app

COPY . .
RUN go mod download
RUN go build -ldflags="-X 'main.version=$VERSION'" -v server.go

FROM registry.trendyol.com/platform/base/image/alpine:3.10.1 AS alpine

ENV LANG C.UTF-8

ENV PROXY_URL ""
ENV COUCHBASE_HOST ""
ENV COUCHBASE_USERNAME ""
ENV COUCHBASE_PASSWORD ""
ENV BUCKET_NAME ""

RUN apk --no-cache add tzdata ca-certificates
COPY --from=builder /app/main   /app/server

WORKDIR /app

RUN chmod +x server

EXPOSE 3083

ENTRYPOINT ["./main","server"]
