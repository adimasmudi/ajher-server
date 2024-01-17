# Use the official Golang image as the base image
FROM golang:1.17 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download and install dependencies
RUN go mod download

# Copy the entire application source code into the container
COPY . .

# Build the Go application with debugging information
RUN CGO_ENABLED=0 GOOS=linux go build -o app -gcflags "all=-N -l" .

# Use a smaller base image for the final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/app .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the executable
CMD ["./app"]
