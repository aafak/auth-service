# Dockerfile
FROM golang:1.20-alpine AS builder

ARG PROXY

# Set environment variables for the proxy
ENV http_proxy=${PROXY}
ENV https_proxy=${PROXY}
ENV no_proxy=localhost,127.0.0.1

# install build tools
RUN set -eux; \
    apk add -U --no-cache \
        curl \
        git  \
        make \
        bash \
    ;

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN ls

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/

FROM alpine:latest  

# RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .
# COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./main"]