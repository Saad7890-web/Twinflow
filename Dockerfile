
FROM golang:1.24-alpine AS builder

WORKDIR /app


RUN apk add --no-cache git


COPY go.mod go.sum ./
RUN go mod download


COPY . .


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o twinflow ./cmd/twinflow


FROM alpine:latest

WORKDIR /root/


COPY --from=builder /app/twinflow .


RUN mkdir -p /captures

# Expose default port (proxy)
EXPOSE 8080


ENTRYPOINT ["./twinflow"]