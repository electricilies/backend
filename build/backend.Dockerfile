FROM golang:1.24.5-alpine3.21 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o /backend ./cmd/main.go

FROM scratch
COPY --from=builder /backend /backend
EXPOSE 8080
ENTRYPOINT ["/backend"]

