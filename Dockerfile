# Build stage
FROM golang:1.19-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o ./server ./cmd/server

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
COPY example.env .
COPY wait-for.sh .
ENV PORT=8080
EXPOSE $PORT

CMD ["/app/server"]

