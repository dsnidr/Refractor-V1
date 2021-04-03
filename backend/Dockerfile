FROM golang:alpine AS build

# Install build tools
RUN apk --no-cache add gcc g++ make git

# Set build env variables
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

COPY . .

RUN go build -ldflags "-s -w" -o refractor-bin -i cmd/refractor/main.go

# Create actual container
FROM alpine

RUN apk --no-cache add ca-certificates

WORKDIR /var/refractor

# Copy the binary from the build stage into /var/refractor
COPY --from=build /build/refractor-bin ./refractor

ENTRYPOINT PORT=80 /var/refractor/refractor
