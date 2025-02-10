# Use a minimal base image
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Copy the Go code and build
COPY main.go .
RUN go mod init logoutput && go mod tidy && go build -o logoutput

# Use a small base image for final execution
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/logoutput .

# Run the application
CMD ["./logoutput"]
