
FROM golang as builder

RUN go get github.com/kaishuu0123/cors-reverse-proxy

FROM alpine

WORKDIR /
COPY --from=builder /go/bin/cors-reverse-proxy .

ENTRYPOINT ["/cors-reverse-proxy"]