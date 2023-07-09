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
ENV PORT=8080
EXPOSE $PORT

# 加载环境变量
RUN apk add --no-cache bash
CMD [ "bash", "-c", "source ./example.env && ./server" ]

