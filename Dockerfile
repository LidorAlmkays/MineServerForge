# Stage 1: Build stage
FROM golang:latest AS build

# Set the working directory
WORKDIR /app

# Copy the source code
COPY ./apps/frontend_gateway .

RUN go mod download

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./cmd/main.go

# Stage 2: Final stage
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/myapp .
COPY --from=build /app/configs ./configs

# Set the entrypoint command
ENTRYPOINT ["/app/myapp", "--Mode","production"]