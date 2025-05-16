FROM golang:1.24-alpine AS builder

WORKDIR /app

# Download dependencies first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest and build
COPY . .
RUN cd cmd && go build -o ../main .

# Final stage
FROM alpine:latest AS final

WORKDIR /app

COPY --from=builder /app/main .
COPY .env .


# Install CA certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

EXPOSE 8080

CMD ["./main"]