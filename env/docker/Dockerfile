# Use newer version of Golang and Alpine Linux as base image
FROM golang:1.22-alpine AS builder

# Update packages and install needed dependencies
RUN apk update && \
    apk add ca-certificates git curl gcc musl-dev build-base autoconf automake libtool

# Install air for hot reloading of Go applications during development
RUN GO111MODULE=on go install github.com/cosmtrek/air@latest

# Set working directory
WORKDIR /app

# Copy the source code
COPY . .

# Set the command that will be ran when container is started
CMD ["air"]
