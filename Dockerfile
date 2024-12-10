FROM golang:alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o my-go-api

FROM alpine
WORKDIR /app
COPY --from=builder /app/my-go-api .
CMD ["./my-go-api"]
