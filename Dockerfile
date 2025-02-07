FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build

FROM alpine

RUN apk --no-cache add ca-certificates

WORKDIR /
COPY --from=builder /app/cors-reverse-proxy .

ENTRYPOINT ["/cors-reverse-proxy"]