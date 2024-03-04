# Start from a base image containing the Go runtime
FROM golang:1.22.0 AS build


ARG Service_Name
# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY ./go.mod .
COPY ./go.sum .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 go build -o main ./{Service_Name}/cmd/main.go 

FROM alpine:3.19.1 AS production


COPY --from=build /app/main .

EXPOSE 8080
ENTRYPOINT [ "./main" ]