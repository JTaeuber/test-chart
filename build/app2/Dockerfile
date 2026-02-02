# Build stage
FROM golang:1.25.6-alpine AS builder
WORKDIR /tmp/app
COPY ../../go.mod /tmp/app/go.mod
RUN go mod download
COPY ../.. /tmp/app
RUN CGO_ENABLED=0 go build -o bin/app ./test-application/test-application.go

# Runtime stage
FROM gcr.io/distroless/base-debian13
COPY --from=builder /tmp/app/bin/app /app/app2
ENTRYPOINT ["/app/app2"]
