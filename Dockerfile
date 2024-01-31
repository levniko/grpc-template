# Start from the official Golang image to create a build artifact.
FROM golang:1.21.6 as builder

# Set the Current Working Directory inside the container.
WORKDIR /app

# Copy go.mod and go.sum files to the workspace.
COPY go.* ./

# Download all dependencies.
RUN go mod download

# Copy the source code into the container.
COPY . .

# Build the application.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch for the final image.
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage.
COPY --from=builder /app/main .

# Command to run the executable.
CMD ["./cmd/main"]
