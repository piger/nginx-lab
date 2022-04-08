# First stage
FROM golang:1.18 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -v -o backend

# Second stage
FROM debian:stable-slim

ARG DEBIAN_FRONTEND=noninteractive

RUN apt update && \
    apt upgrade -y && \
    apt install -y --no-install-recommends ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/backend /app/
CMD ["/app/backend"]
