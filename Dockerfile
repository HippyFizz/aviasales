FROM golang:1.12-stretch as builder

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

WORKDIR /aviasales

ADD . .

ENV CGO_ENABLED=0
ENV GOOS=linux
RUN make deps build

FROM alpine:3.9

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN apk add --no-cache ca-certificates

COPY --from=builder /aviasales/aviasales /go/bin/aviasales
COPY --from=builder /aviasales/config/config.yaml $GOPATH/config/config.yaml

RUN mkdir -p /go/var /go/var/run /go/var/log /go/var/tmp

EXPOSE 7000

VOLUME /go/var

ENTRYPOINT ["/go/bin/aviasales"]
