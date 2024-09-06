# build the project with golang image
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

# RUN go mod download OR go mod tidy
# tidy will automatically download ONLY required dependencies
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/api

# Build the final image
FROM ubuntu:22.04
EXPOSE 80
COPY --from=builder /app/server app
RUN chmod +x /app
CMD ["./app"]