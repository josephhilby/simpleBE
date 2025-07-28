# Use a lightweight Go image with Alpine
FROM golang:1.23-alpine as builder

# Install build dependencies
RUN apk add --no-cache git

WORKDIR /app

# Copy Go mod files first to cache dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy full source and build the app
COPY . ./
RUN go build -o server ./cmd

# Stage 2: Use a minimal runtime (no shell, secure)
FROM gcr.io/distroless/base-debian11

WORKDIR /app
COPY --from=builder /app/server .
CMD ["./server"]