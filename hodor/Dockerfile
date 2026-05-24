FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go build -o hodor main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates wget
WORKDIR /app
COPY --from=builder /app/hodor .
EXPOSE 8080
CMD ["./hodor"]