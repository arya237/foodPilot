# Stage 1: Build
FROM golang:1.24-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o FoodPilot ./main.go

# Stage 2: Run 
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/FoodPilot .

EXPOSE 8080

ENTRYPOINT ["./FoodPilot"]
