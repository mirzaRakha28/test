# Dockerfile

# Use the appropriate base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o /docker-go-app

# Expose port 6060
EXPOSE 6060

# Run the application
CMD ["/docker-go-app"]
