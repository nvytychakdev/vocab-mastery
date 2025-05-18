# Start with the official Golang image as a build stage
FROM golang:1.24.3-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./vocab-mastery -v ./cmd/app/main.go 



# Second stage to run the server
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/vocab-mastery .

# Expose port (optional, match your app's port)
EXPOSE 8080

# Command to run the application
CMD ["./vocab-mastery"]