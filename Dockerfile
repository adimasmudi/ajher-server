# Use a more specific version of the Go base image
FROM golang:1.19.10-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy only the go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download and install dependencies
RUN go mod download

# Copy the entire application source code into the container
COPY . .

EXPOSE 5000

# Build the Go application
RUN go build -o ./out/dist .

# Set the command to run the executable
CMD ["./out/dist"]