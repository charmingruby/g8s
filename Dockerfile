# Build Stage
FROM golang:1.22-alpine AS builder
WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/g8s ./cmd/api/main.go

# Run Stage
FROM alpine:latest
WORKDIR /app
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk*
COPY --from=builder /build/bin/g8s ./g8s
COPY .env /app
EXPOSE 3333
CMD [ "/app/g8s" ]