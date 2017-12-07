FROM golang:1.9-alpine AS builder

RUN apk add --no-cache git build-base

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/aptible/mini-collector/

COPY Gopkg.lock .
COPY Gopkg.toml .
RUN dep ensure -vendor-only

COPY . .
RUN go build -i github.com/aptible/mini-collector/cmd/mini-collector
RUN go build -i github.com/aptible/mini-collector/cmd/aggregator


FROM alpine:3.7
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/aptible/mini-collector/mini-collector /usr/local/bin/
CMD ["mini-collector"]
