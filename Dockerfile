# Use the official Golang image
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main .

# Expose port
EXPOSE 8080

# Run the executable
CMD ["./main"]