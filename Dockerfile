FROM golang:alpine AS builder
WORKDIR /app
COPY proto /app/proto
COPY . .
EXPOSE 8085
ENTRYPOINT ["go", "run", "main.go"]