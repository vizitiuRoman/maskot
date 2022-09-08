FROM golang:alpine as base

# ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Vizitiu Roman vizitiuroman.md@gmail.com"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container
WORKDIR /app

FROM base as application-dev

RUN apk --no-cache add entr
RUN apk --no-cache add npm

RUN npm install -g nodemon

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY .. .

RUN echo " \
    (while true; do nodemon --exec go run cmd/main.go --signal SIGTERM || sleep 10; done) \
    " > /entrypoint.sh \
    && chmod +x /entrypoint.sh

ENTRYPOINT /entrypoint.sh

FROM base as build

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY .. .
# Build the Go app
RUN cd cmd && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest as application
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=build /app/cmd/main .

# Expose port 8080 to the outside world
EXPOSE 8091

#Command to run the executable
CMD ["./main"]