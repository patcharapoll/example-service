# Start from golang base image
FROM golang:1.16.2-alpine as dependencies

ENV GO11MODULE=on

# Install git.
# Git is required ofr fetching the dependencies.
RUN apk update && apk add --no-cache git make gcc libc-dev protobuf-dev

RUN go get github.com/golang/protobuf/protoc-gen-go
RUN go get github.com/favadi/protoc-go-inject-tag

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN make pbgen
RUN make build

# Start a new stage from scratch
# FROM scratch
FROM alpine

# # Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=dependencies /app/bin /app/bin
COPY --from=dependencies /app/entrypoint.sh /

RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]