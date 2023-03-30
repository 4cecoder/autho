# Build stage
FROM golang:1.17-alpine as builder

# Install git and ca-certificates
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o autho

# Final stage
FROM scratch

# Copy necessary files from builder stage
COPY --from=builder /app/autho /app/autho
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Set environment variables
ENV APP_PORT=8080
ENV DB_USERNAME=${DB_USERNAME}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_HOST=${DB_HOST}
ENV DB_NAME=${DB_NAME}

# Expose the application port
EXPOSE 8080

# Run the application
ENTRYPOINT ["/app/autho"]
