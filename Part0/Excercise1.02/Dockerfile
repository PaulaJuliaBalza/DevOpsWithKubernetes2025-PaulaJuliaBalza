# Use a minimal Go image
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY main.go .
RUN go mod init todoapp && go mod tidy && go build -o todoapp

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/todoapp .

EXPOSE 8080
CMD ["./todoapp"]
