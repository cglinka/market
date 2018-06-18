FROM golang:1.10.2

ADD . /go/src/github.com/cglinka/market
RUN go install github.com/cglinka/market

ENTRYPOINT /go/bin/market