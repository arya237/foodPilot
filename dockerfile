# Stage 1: Build
FROM golang:1.24-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.io

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o FoodPilot ./main.go

# Stage 2: Run 
FROM alpine:latest

WORKDIR /app

# Copy the binary
COPY --from=builder /app/FoodPilot .

# Copy static files (CSS, JS)
COPY --from=builder /app/statics ./statics

# Copy HTML templates
COPY --from=builder /app/internal/delivery/web/ui/templates ./internal/delivery/web/ui/templates

EXPOSE 8080

ENTRYPOINT ["./FoodPilot"]
