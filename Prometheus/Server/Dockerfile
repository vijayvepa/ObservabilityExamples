# Build Stage
FROM golang:1.24.3 AS builder
WORKDIR /workspace
COPY go.mod  go.mod
COPY go.sum  go.sum
RUN go mod download

# Copy the go source
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o server main.go

# Publish Stage
FROM alpine:latest

RUN apk update
RUN apk add --no-cache bash curl jq
WORKDIR /
COPY --from=builder /workspace/server .
RUN chmod +x server
ENTRYPOINT [ "./server" ]
EXPOSE 8090