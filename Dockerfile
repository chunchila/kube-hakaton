## Build stage
FROM golang:1.9.4 as builder

# Set the working directory to the app directory
WORKDIR /go/src/app



# Copy the application files
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

## App stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/app .



# Expose the application on port 8081
EXPOSE 8081

# Set the entry point of the container to the bee command that runs the
# application and watches for changes
CMD ["./app", "run"]