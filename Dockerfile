FROM golang:1.24-alpine AS builder

WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download


COPY . .
RUN cd cmd && go build -o ../main .


FROM alpine:latest AS final

WORKDIR /app

COPY --from=builder /app/main .
COPY .env .



RUN apk --no-cache add ca-certificates

EXPOSE 8080

CMD ["./main"]