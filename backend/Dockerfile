# First stage
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o backend .

# Second stage
FROM scratch
COPY --from=builder /app/backend /app/
CMD ["/app/backend"]
